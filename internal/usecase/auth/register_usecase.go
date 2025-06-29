package authUC

import (
	"cms-server/constants"
	"cms-server/internal/entity"
	"cms-server/internal/repository"
	serviceJwt "cms-server/internal/service/jwt"
	"cms-server/internal/service/queue"
	"context"
	"errors"
	"log"
	"runtime"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/diebietse/gotp/v2"
	"github.com/go-pg/pg/v10"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type ResRegister struct {
	UserInfor entity.UserInfor
	MailTpl   *entity.MailTemplate
	Token     string
}

type RegisterReq struct {
	Email           string
	FullName        string
	Password        string
	ConfirmPassword string
	Code            string
}

type RegisterUsecase interface {
	CheckUserExist(email string) (entity.User, error)
	hashPassword(password string) (string, error)
	Register(user RegisterReq, os string, exp time.Time) (ResRegister, error)
	GengerateCode(time time.Time, secrest string) (string, error)
	createOrUpdateUser(user RegisterReq, ctx context.Context) (entity.UserInfor, *entity.MailTemplate, error)
	saveToken(token string, id string, os string) error
	SendMail(tlp *entity.MailTemplate, user entity.UserInfor, linkVerify string) error
}

type registerUsecaseImpl struct {
	userRepo          repository.UserRepository
	mailTplRepo       repository.MailTemplateRepository
	mailHistoryRepo   repository.MailHistoryRepository
	statusHistoryRepo repository.StatusHistoryRepository
	sessionRepo       repository.SessionRepository
	jwt               serviceJwt.JwtService
	qc                queue.QueueClient
	tx                repository.ManagerTransaction
}

func NewRegisterUsecase(
	userRepo repository.UserRepository,
	mailTplRepo repository.MailTemplateRepository,
	mailHistoryRepo repository.MailHistoryRepository,
	statusHistoryRepo repository.StatusHistoryRepository,
	sessionRepo repository.SessionRepository,
	jwt serviceJwt.JwtService,
	qc queue.QueueClient,
	tx repository.ManagerTransaction,
) RegisterUsecase {
	return &registerUsecaseImpl{
		userRepo:          userRepo,
		mailHistoryRepo:   mailHistoryRepo,
		mailTplRepo:       mailTplRepo,
		statusHistoryRepo: statusHistoryRepo,
		sessionRepo:       sessionRepo,
		qc:                qc,
		tx:                tx,
		jwt:               jwt,
	}
}

func (uc *registerUsecaseImpl) CheckUserExist(email string) (entity.User, error) {
	return uc.userRepo.GetUserByEmail(email)
}
func (uc *registerUsecaseImpl) GengerateCode(time time.Time, secrest string) (string, error) {
	secret, _ := gotp.DecodeBase32(secrest)
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

func (uc *registerUsecaseImpl) Register(user RegisterReq, os string, exp time.Time) (ResRegister, error) {
	res := ResRegister{}
	err := uc.tx.Do(func(ctx context.Context) error {
		var err error
		if res.UserInfor, res.MailTpl, err = uc.createOrUpdateUser(user, ctx); err != nil {
			log.Println("Error creating or updating user:", err)
			return err
		}
		if err = uc.sessionRepo.DeleteSessionVerifyByUserID(res.UserInfor.ID); err != nil {
			return err
		}
		if res.Token, err = uc.jwt.GenRegisterToken(res.UserInfor.ID, user.Code, exp); err != nil {
			return err
		}
		if err = uc.saveToken(res.Token, res.UserInfor.ID, os); err != nil {
			return err
		}
		return nil
	})

	// err := uc.tx.RunInTransaction(func(tx *pg.Tx) error {
	// 	var err error
	// 	if res.UserInfor, res.MailTpl, err = uc.createOrUpdateUser(user, tx); err != nil {
	// 		return err
	// 	}
	// 	if err = uc.sessionRepo.DeleteSessionVerifyByUserID(res.UserInfor.ID); err != nil {
	// 		return err
	// 	}
	// 	if res.Token, err = uc.jwt.GenRegisterToken(res.UserInfor.ID, user.Code, exp); err != nil {
	// 		return err
	// 	}
	// 	if err = uc.saveToken(res.Token, res.UserInfor.ID, os); err != nil {
	// 		return err
	// 	}
	// 	return nil
	// })

	return res, err
}

func (uc *registerUsecaseImpl) createOrUpdateUser(user RegisterReq, ctx context.Context) (entity.UserInfor, *entity.MailTemplate, error) {
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

	if isExist, err := uc.userRepo.CheckUserExist(newUser.Email); err != nil {
		return userInfo, tpl, err
	} else if !isExist {
		log.Println("Tạo mới người dùng:", newUser.Email)
		if userInfo, err = uc.userRepo.Tx(ctx).CreateUser(newUser); err != nil {
			return userInfo, tpl, err
		}
	} else {
		newUser.ID = "" // Nó sẽ không cập nhập ID bởi ID là khóa chính | thêm cho dễ hiểu
		if ok, err := uc.userRepo.Tx(ctx).UpdateUserByEmail(newUser.Email, newUser); err != nil {
			return userInfo, tpl, err
		} else if u, err := uc.userRepo.GetUserByEmail(newUser.Email); ok && err == nil {
			userInfo = u.GetInfor()
		} else {
			return userInfo, tpl, err
		}
	}
	if tpl, err = uc.mailTplRepo.GetMailTplById(constants.TPL_REGISTER_MAIL); err != nil {
		return userInfo, tpl, err
	}

	return userInfo, tpl, errors.New("không tìm thấy mẫu email đăng ký")
}

func (uc *registerUsecaseImpl) saveToken(token string, userId string, os string) error {
	return uc.sessionRepo.CreateSession(entity.Session{
		Token:     token,
		UserID:    userId,
		Os:        os,
		Type:      entity.SessionTypeVerify,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(constants.VerifyExpiredAt * time.Second),
	})
}

func (uc *registerUsecaseImpl) SendMail(tlp *entity.MailTemplate, user entity.UserInfor, linkVerify string) error {
	data := map[string]any{
		"link": linkVerify,
		"user": user,
	}
	payload := queue.Payload{
		Provider: tlp.ProviderEmail,
		Template: tlp.ID,
		Data:     data,
		To:       &user.Email,
	}
	Id, err := uc.qc.EnqueueMail(payload)
	if err != nil {
		return err
	}

	return uc.tx.RunInTransaction(func(tx *pg.Tx) error {
		err := uc.mailHistoryRepo.Create(&entity.MailHistory{
			ID:            Id,
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
			MailHistoryId: Id,
			Message:       "Send mail pending",
			CreatedAt:     time.Now(),
		})
		return err
	})
}
