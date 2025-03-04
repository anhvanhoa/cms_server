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
	hashPassword(password string) (string, error)
	CreateUser(user modelauth.RegisterReq) (entity.UserInfor, *entity.MailTemplate, error)
	SendMail(tpl *entity.MailTemplate, code string, user entity.UserInfor) error
}

type registerUsecaseImpl struct {
	userRepo    repository.UserRepository
	mailTplRepo repository.MailTemplateRepository
	qc          bootstrap.QueueClient
}

func NewRegisterUsecase(userRepo repository.UserRepository, mailTplRepo repository.MailTemplateRepository, qc bootstrap.QueueClient) RegisterUsecase {
	return &registerUsecaseImpl{
		userRepo:    userRepo,
		qc:          qc,
		mailTplRepo: mailTplRepo,
	}
}

func (uc *registerUsecaseImpl) CheckUserExist(email string) (bool, error) {
	return uc.userRepo.CheckUserExist(email)
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
		ID:       id,
		Email:    user.Email,
		Password: user.Password,
		FullName: user.FullName,
	}

	if newUser.Password, err = uc.hashPassword(newUser.Password); err != nil {
		return userInfo, tpl, err
	}

	uc.userRepo.RunInTransaction(func(tx *pg.Tx) error {
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

func (uc *registerUsecaseImpl) SendMail(tlp *entity.MailTemplate, code string, user entity.UserInfor) error {
	task, err := uc.qc.NewTaskMailSystem(map[string]any{
		"code": code,
		"user": user,
	})
	if err != nil {
		return err
	}
	if info, err := uc.qc.Enqueue(task); err != nil {
		return err
	} else {
		fmt.Println("info send: ", info)
	}
	return nil
}
