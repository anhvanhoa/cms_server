package worker

import (
	"cms-server/bootstrap"
	"cms-server/constants"
	"cms-server/internal/entity"
	"cms-server/internal/repository"
	pkgerror "cms-server/pkg/error"
	pkglog "cms-server/pkg/logger"
	"cms-server/pkg/mailtemplate"
	"context"
	"crypto/tls"
	"encoding/json"
	"time"

	"github.com/hibiken/asynq"
)

type EmailSystem struct {
	log               pkglog.Logger
	mailTemplate      mailtemplate.MailTemplate
	mailProvider      bootstrap.MailProvider
	mailTplRepo       repository.MailTemplateRepository
	mailProvierRepo   repository.MailProviderRepository
	mailHistoryRepo   repository.MailHistoryRepository
	statusHistoryRepo repository.StatusHistoryRepository
}

func (e *EmailSystem) ProcessTask(ctx context.Context, task *asynq.Task) error {
	var payload bootstrap.Payload
	statusErr := entity.StatusHistory{
		Status:        entity.MAIL_STATUS_FAILED,
		MailHistoryId: task.ResultWriter().TaskID(),
		CreatedAt:     time.Now(),
	}
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		statusErr.Message = "Failed to parse payload: " + err.Error()
		e.statusHistoryRepo.Create(&statusErr)
		e.log.Warn(statusErr.Message)
		return pkgerror.New(statusErr.Message)
	}
	tpl, err := e.mailTplRepo.GetMailTplById(payload.Template)
	if err != nil {
		statusErr.Message = "Failed to get mail template: " + err.Error()
		e.statusHistoryRepo.Create(&statusErr)
		return err
	} else if tpl == nil {
		statusErr.Message = "Template not found"
		e.statusHistoryRepo.Create(&statusErr)
		return pkgerror.New("Không tìm thấy mẫu email")
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
		return pkgerror.New("Không tìm thấy cấu hình gửi email")
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
	// Gửi email
	if err := e.mailProvider.SendMail(tos, mailT.Subject, mailT.Body, payload.Data); err != nil {
		statusErr.Message = "Failed to send mail: " + err.Error()
		e.statusHistoryRepo.Create(&statusErr)
		return err
	}

	// Update lại subject và body vào mail history
	err = e.mailHistoryRepo.UpdateSubAndBodyById(task.ResultWriter().TaskID(), mailT.Subject, mailT.Body)
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

func NewEmailSystem(
	mux *asynq.ServeMux,
	log pkglog.Logger,
	mailtemplate mailtemplate.MailTemplate,
	mailProvider bootstrap.MailProvider,
	mailTplRepo repository.MailTemplateRepository,
	mailProvierRepo repository.MailProviderRepository,
	mailHistoryRepo repository.MailHistoryRepository,
	statusHistoryRepo repository.StatusHistoryRepository,
) {
	mux.Handle(string(constants.QUEUE_EMAIL_SYSTEM), &EmailSystem{
		log:               log,
		mailTemplate:      mailtemplate,
		mailProvider:      mailProvider,
		mailTplRepo:       mailTplRepo,
		mailProvierRepo:   mailProvierRepo,
		mailHistoryRepo:   mailHistoryRepo,
		statusHistoryRepo: statusHistoryRepo,
	})
}
