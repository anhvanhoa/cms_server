package bootstrap

import (
	"cms-server/constants"
	"encoding/json"
	"time"

	"github.com/hibiken/asynq"
)

type To string

type Payload struct {
	Provider  string
	Tos       *[]To
	To        *To
	Templates string
	Data      map[string]any
}

type QueueClient interface {
	NewTask(typeTask string, payload map[string]any, opts ...asynq.Option) (*asynq.Task, error)
	NewTaskMailSystem(payload map[string]any, opts ...asynq.Option) (*asynq.Task, error)
	NewTaskSms(payload map[string]any, opts ...asynq.Option) (*asynq.Task, error)
	Enqueue(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error)
	Ping() error
	Close()
}

type queueClient struct {
	client  *asynq.Client
	retry   int
	timeout time.Duration
}

// NewQueueClient creates a new queue client
func NewQueueClient(env *Env) QueueClient {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     env.QUEUE.Addr,
		DB:       env.QUEUE.DB,
		Password: env.QUEUE.Password,
		Network:  env.QUEUE.Network,
	})
	return &queueClient{
		client:  client,
		retry:   5,
		timeout: 10 * time.Minute,
	}
}

func (qc *queueClient) Close() {
	qc.client.Close()
}

func (qc *queueClient) NewTask(typeTask string, payload map[string]any, opts ...asynq.Option) (*asynq.Task, error) {
	defaultOpts := []asynq.Option{
		asynq.MaxRetry(qc.retry),
		asynq.Timeout(qc.timeout),
	}
	opts = append(defaultOpts, opts...)
	pl, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(typeTask, pl, opts...), nil
}

func (qc *queueClient) NewTaskMailSystem(payload map[string]any, opts ...asynq.Option) (*asynq.Task, error) {
	return qc.NewTask(string(constants.QUEUE_EMAIL_SYSTEM), payload, opts...)
}

func (qc *queueClient) NewTaskSms(payload map[string]any, opts ...asynq.Option) (*asynq.Task, error) {
	return qc.NewTask(string(constants.QUEUE_SMS), payload, opts...)
}

func (qc *queueClient) Enqueue(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	return qc.client.Enqueue(task, opts...)
}

func (qc *queueClient) Ping() error {
	return qc.client.Ping()
}
