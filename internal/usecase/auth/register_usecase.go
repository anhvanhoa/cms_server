package auth

import (
	"cms-server/bootstrap"
	"cms-server/constants"
	"cms-server/internal/entity"
	modelauth "cms-server/internal/model/auth"
	"cms-server/internal/repository"
	pkgjwt "cms-server/pkg/jwt"
	"runtime"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/diebietse/gotp/v2"
	"github.com/go-pg/pg/v10"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type RegisterUsecase interface {
	CheckUserExist(email string) (bool, error)
	hashPassword(password string) (string, error)
	GengerateCode(time time.Time) (string, error)
	CreateUser(user modelauth.RegisterReq) (entity.UserInfor, *entity.MailTemplate, error)
	GengerateToken(data pkgjwt.RegisterClaims) (string, error)
	SendMail(tpl *entity.MailTemplate, token string, user entity.UserInfor) error
}

type registerUsecaseImpl struct {
	userRepo          repository.UserRepository
	mailTplRepo       repository.MailTemplateRepository
	mailHistoryRepo   repository.MailHistoryRepository
	statusHistoryRepo repository.StatusHistoryRepository
	jwt               pkgjwt.JWT
	qc                bootstrap.QueueClient
	tx                repository.ManagerTransaction
	env               *bootstrap.Env
}

func NewRegisterUsecase(
	userRepo repository.UserRepository,
	mailTplRepo repository.MailTemplateRepository,
	mailHistoryRepo repository.MailHistoryRepository,
	statusHistoryRepo repository.StatusHistoryRepository,
	jwt pkgjwt.JWT,
	qc bootstrap.QueueClient,
	tx repository.ManagerTransaction,
	env *bootstrap.Env,
) RegisterUsecase {
	return &registerUsecaseImpl{
		userRepo:          userRepo,
		mailHistoryRepo:   mailHistoryRepo,
		mailTplRepo:       mailTplRepo,
		statusHistoryRepo: statusHistoryRepo,
		qc:                qc,
		tx:                tx,
		env:               env,
		jwt:               jwt,
	}
}

func (uc *registerUsecaseImpl) CheckUserExist(email string) (bool, error) {
	return uc.userRepo.CheckUserExist(email)
}
func (uc *registerUsecaseImpl) GengerateCode(time time.Time) (string, error) {
	secret, _ := gotp.DecodeBase32(uc.env.SECRET_OTP)
	hotp, _ := gotp.NewHOTP(secret)
	return hotp.At(int(time.Unix()))
}
func (uc *registerUsecaseImpl) hashPassword(password string) (string, error) {
	params := argon2id.Params{
		Memory:      64 * 1024,
		Iterations:  4,
		Parallelism: uint8(runtime.NumCPU()),
		SaltLength:  16,
		KeyLength:   32,
	}
	return argon2id.CreateHash(password, &params)
}
func (uc *registerUsecaseImpl) CreateUser(user modelauth.RegisterReq) (entity.UserInfor, *entity.MailTemplate, error) {
	var userInfo entity.UserInfor
	var tpl *entity.MailTemplate
	id, err := gonanoid.New(10)
	if err != nil {
		return userInfo, tpl, err
	}
	newUser := entity.User{
		ID:         id,
		Email:      user.Email,
		Password:   user.Password,
		FullName:   user.FullName,
		CodeVerify: user.Code,
	}

	if newUser.Password, err = uc.hashPassword(newUser.Password); err != nil {
		return userInfo, tpl, err
	}

	err = uc.tx.RunInTransaction(func(tx *pg.Tx) error {
		if userInfo, err = uc.userRepo.CreateUser(newUser, tx); err != nil {
			return err
		}
		if tpl, err = uc.mailTplRepo.GetMailTplById(constants.TPL_REGISTER_MAIL); err != nil {
			return err
		}
		return nil
	})
	return userInfo, tpl, err
}
func (uc *registerUsecaseImpl) GengerateToken(data pkgjwt.RegisterClaims) (string, error) {
	return uc.jwt.GenRegisterToken(data)
}
func (uc *registerUsecaseImpl) SendMail(tlp *entity.MailTemplate, token string, user entity.UserInfor) error {
	link := uc.env.FRONTEND_URL + "/verify/" + token
	data := map[string]any{
		"link": link,
		"user": user,
	}
	task, err := uc.qc.NewTaskMailSystem(bootstrap.Payload{
		Provider: tlp.ProviderEmail,
		Template: tlp.ID,
		Data:     data,
		To:       &user.Email,
	})
	if err != nil {
		return err
	}
	if info, err := uc.qc.Enqueue(task); err != nil {
		return err
	} else {
		return uc.tx.RunInTransaction(func(tx *pg.Tx) error {
			err := uc.mailHistoryRepo.Create(&entity.MailHistory{
				ID:            info.ID,
				TemplateId:    tlp.ID,
				To:            user.Email,
				Data:          data,
				EmailProvider: tlp.ProviderEmail,
			})
			if err != nil {
				return err
			}
			err = uc.statusHistoryRepo.Create(&entity.StatusHistory{
				Status:        entity.MAIL_STATUS_PENDING,
				MailHistoryId: info.ID,
				Message:       "Send mail pending",
				CreatedAt:     time.Now(),
			})
			return err
		})
	}
}
