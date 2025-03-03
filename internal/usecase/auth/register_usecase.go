package auth

import (
	"cms-server/bootstrap"
	"cms-server/constants"
	"cms-server/internal/entity"
	modelauth "cms-server/internal/model/auth"
	"cms-server/internal/repository"
	"fmt"
	"runtime"

	"github.com/alexedwards/argon2id"
	"github.com/go-pg/pg/v10"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type RegisterUsecase interface {
	CheckUserExist(email string) (bool, error)
	HashPassword(password string) (string, error)
	CreateUser(user modelauth.RegisterReq) (entity.UserInfor, error)
	SendMail(code string, user entity.UserInfor) error
}

type registerUsecaseImpl struct {
	userRepo    repository.UserRepository
	mailTplRepo repository.MailTemplateRepository
	qc          bootstrap.QueueClient
	tm          *bootstrap.TransactionManager
}

func NewRegisterUsecase(userRepo repository.UserRepository, mailTplRepo repository.MailTemplateRepository, qc bootstrap.QueueClient, tm *bootstrap.TransactionManager) RegisterUsecase {
	return &registerUsecaseImpl{
		userRepo:    userRepo,
		qc:          qc,
		mailTplRepo: mailTplRepo,
		tm:          tm,
	}
}

func (uc *registerUsecaseImpl) CheckUserExist(email string) (bool, error) {
	return uc.userRepo.CheckUserExist(email)
}

func (uc *registerUsecaseImpl) HashPassword(password string) (string, error) {
	params := argon2id.Params{
		Memory:      64 * 1024,
		Iterations:  4,
		Parallelism: uint8(runtime.NumCPU()),
		SaltLength:  16,
		KeyLength:   32,
	}
	return argon2id.CreateHash(password, &params)
}
func (uc *registerUsecaseImpl) CreateUser(user modelauth.RegisterReq) (entity.UserInfor, error) {
	id, err := gonanoid.New(10)
	if err != nil {
		return entity.UserInfor{}, err
	}
	newUser := entity.User{
		ID:       id,
		Email:    user.Email,
		Password: user.Password,
		FullName: user.FullName,
	}

	var userInfo entity.UserInfor
	err = uc.tm.WithTransaction(func(tx *pg.Tx) error {
		// _, err := uc.userRepo.CreateUserTx(tx, newUser)
		userInfo = newUser.GetInfor()
		return err
	})

	return userInfo, err
}

func (uc *registerUsecaseImpl) SendMail(code string, user entity.UserInfor) error {
	tpl, err := uc.mailTplRepo.GetMailTplById(constants.TPL_REGISTER_MAIL)

	if err != nil {
		return err
	}

	fmt.Println("tpl: ", tpl)

	// task, err := uc.qc.NewTaskMailSystem(map[string]any{
	// 	"code": code,
	// 	"user": user,
	// })
	// if err != nil {
	// 	return err
	// }
	// if info, err := uc.qc.Enqueue(task); err != nil {
	// 	return err
	// } else {
	// 	fmt.Println("info send: ", info)
	// }
	return nil
}
