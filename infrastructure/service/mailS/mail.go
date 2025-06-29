package mailS

import (
	"cms-server/bootstrap"
	"cms-server/constants"
	"cms-server/infrastructure/repo"
	"cms-server/infrastructure/service/mailtemplate"
	"cms-server/internal/schedule"
	serviceLogger "cms-server/internal/service/logger"
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/hibiken/asynq"
)

type MailHandler struct {
	mailS schedule.EmailSystemImpl
}

func (e *MailHandler) ProcessTask(ctx context.Context, task *asynq.Task) error {
	return e.mailS.SendMailQueue(task.Payload(), task.ResultWriter().TaskID())
}

func NewEmailHandler(
	mux *asynq.ServeMux,
	env *bootstrap.Env,
	log serviceLogger.Logger,
	db *pg.DB,
) {
	var mailS = schedule.NewEmailSystem(
		log,
		mailtemplate.NewMailTemplate(),
		bootstrap.NewMailProvider(),
		repo.NewMailTplRepository(db),
		repo.NewMailProviderRepository(db),
		repo.NewMailHistoryRepository(db),
		repo.NewStatusHistoryRepository(db),
		[]string{env.TEST_EMAIL}, // Danh sách email dùng để test
	)
	mailS.ConfigTest().SetIsProduction(env.IsProduction())
	mux.Handle(string(constants.QUEUE_EMAIL_SYSTEM), &MailHandler{mailS})
}
