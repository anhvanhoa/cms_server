package worker

import (
	"cms-server/constants"
	pkglog "cms-server/pkg/logger"
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

type EmailSystem struct {
	log pkglog.Logger
}

func (e *EmailSystem) ProcessTask(ctx context.Context, task *asynq.Task) error {
	var payload map[string]any
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		e.log.Warn("failed to parse payload: " + err.Error())
		return fmt.Errorf("failed to parse payload: %w", err)
	}
	fmt.Println("EmailSystem task is running")
	fmt.Println(payload)
	// Xử lý gửi email
	fmt.Printf("Sending email to %s with subject %s\n", payload["to"], payload["subject"])
	return nil
}

func NewEmailSystem(mux *asynq.ServeMux, log pkglog.Logger) {
	mux.Handle(string(constants.QUEUE_EMAIL_SYSTEM), &EmailSystem{
		log: log,
	})
}
