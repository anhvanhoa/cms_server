package schedule

import (
	"cms-server/bootstrap"
	"cms-server/domain/entity"
	"cms-server/domain/repository"
	serviceLogger "cms-server/domain/service/logger"
	mailtpl "cms-server/domain/service/mailTpl"
	"cms-server/domain/service/queue"
	"crypto/tls"
	"encoding/json"
	"errors"
	"time"
)

type EmailSystemImpl interface {
	SendMailQueue(Payload []byte, Id string) error
	ConfigTest() EmailTestingImpl
}

type EmailTestingImpl interface {
	SetIsProduction(mode bool) *EmailTesting
	SetIsAppedMail(mode bool) *EmailTesting
	SetTestMails(mails []string) *EmailTesting
}

type EmailTesting struct {
	TestMails    []string // Danh sách email dùng để test
	IsAppedMail  bool
	IsProduction bool // Biến để xác định môi trường
}

type EmailSystem struct {
	configTest        EmailTesting
	log               serviceLogger.Logger
	mailTemplate      mailtpl.MailTemplate
	mailProvider      bootstrap.MailProvider
	mailTplRepo       repository.MailTemplateRepository
	mailProvierRepo   repository.MailProviderRepository
	mailHistoryRepo   repository.MailHistoryRepository
	statusHistoryRepo repository.StatusHistoryRepository
}

func (e *EmailSystem) SendMailQueue(Payload []byte, Id string) error {
	var payload queue.Payload
	statusErr := entity.StatusHistory{
		Status:        entity.MAIL_STATUS_FAILED,
		MailHistoryId: Id,
		CreatedAt:     time.Now(),
	}
	if err := json.Unmarshal(Payload, &payload); err != nil {
		statusErr.Message = "Failed to parse payload: " + err.Error()
		e.statusHistoryRepo.Create(&statusErr)
		e.log.Warn(statusErr.Message)
		return errors.New(statusErr.Message)
	}
	tpl, err := e.mailTplRepo.GetMailTplById(payload.Template)
	if err != nil {
		statusErr.Message = "Failed to get mail template: " + err.Error()
		e.statusHistoryRepo.Create(&statusErr)
		return err
	} else if tpl == nil {
		statusErr.Message = "Template not found"
		e.statusHistoryRepo.Create(&statusErr)
		return errors.New("không tìm thấy mẫu email")
	}

	mailT, err := e.mailTemplate.Render(tpl.Subject, tpl.Body, payload.Data)
	if err != nil {
		statusErr.Message = "Failed to render mail template: " + err.Error()
		e.statusHistoryRepo.Create(&statusErr)
		return err
	}

	// Lấy thông tin cấu hình gửi email
	provider, err := e.mailProvierRepo.GetMailProviderByEmail(payload.Provider)
	if err != nil {
		statusErr.Message = "Failed to get mail provider: " + err.Error()
		e.statusHistoryRepo.Create(&statusErr)
		return err
	} else if provider == nil {
		statusErr.Message = "Mail provider not found"
		e.statusHistoryRepo.Create(&statusErr)
		return errors.New("không tìm thấy cấu hình gửi email")
	}
	// Set cấu hình gửi email
	e.mailProvider.SetProvider(&bootstrap.ConfigMail{
		Host:     provider.Host,
		Port:     provider.Port,
		UserName: provider.UserName,
		Password: provider.Password,
		Email:    provider.Email,
		Name:     provider.Name,
		TSL:      &tls.Config{InsecureSkipVerify: true},
	})
	tos := []string{}
	if payload.To != nil {
		tos = append(tos, *payload.To)
	} else if payload.Tos != nil {
		tos = *payload.Tos
	}
	if !e.configTest.IsProduction && len(e.configTest.TestMails) > 0 {
		tos = e.configTest.TestMails
	}
	// Gửi email
	if err := e.mailProvider.SendMail(tos, mailT.Subject, mailT.Body, payload.Data); err != nil {
		statusErr.Message = "Failed to send mail: " + err.Error()
		e.statusHistoryRepo.Create(&statusErr)
		return err
	}

	// Update lại subject và body vào mail history
	err = e.mailHistoryRepo.UpdateSubAndBodyById(Id, mailT.Subject, mailT.Body)
	if err != nil {
		statusErr.Message = "Failed to update mail history: " + err.Error()
		e.statusHistoryRepo.Create(&statusErr)
		return err
	}

	// Thêm trạng thái gửi email thành công
	statusErr.Status = entity.MAIL_STATUS_SENT
	statusErr.Message = "Send mail success"
	return e.statusHistoryRepo.Create(&statusErr)
}

func (e *EmailSystem) ConfigTest() EmailTestingImpl {
	return &e.configTest
}

func NewEmailSystem(
	log serviceLogger.Logger,
	mailtemplate mailtpl.MailTemplate,
	mailProvider bootstrap.MailProvider,
	mailTplRepo repository.MailTemplateRepository,
	mailProvierRepo repository.MailProviderRepository,
	mailHistoryRepo repository.MailHistoryRepository,
	statusHistoryRepo repository.StatusHistoryRepository,
	testMails []string,
) EmailSystemImpl {
	return &EmailSystem{
		configTest:        EmailTesting{TestMails: testMails, IsAppedMail: false, IsProduction: true},
		log:               log,
		mailTemplate:      mailtemplate,
		mailProvider:      mailProvider,
		mailTplRepo:       mailTplRepo,
		mailProvierRepo:   mailProvierRepo,
		mailHistoryRepo:   mailHistoryRepo,
		statusHistoryRepo: statusHistoryRepo,
	}
}

func (et *EmailTesting) SetIsProduction(mode bool) *EmailTesting {
	et.IsProduction = mode
	return et
}

func (et *EmailTesting) SetIsAppedMail(mode bool) *EmailTesting {
	et.IsAppedMail = mode
	return et
}

func (et *EmailTesting) SetTestMails(mails []string) *EmailTesting {
	et.TestMails = mails
	return et
}
