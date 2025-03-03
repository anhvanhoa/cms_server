package main

import (
	"cms-server/bootstrap"
	"cms-server/internal/worker"
	pkglog "cms-server/pkg/logger"

	"github.com/hibiken/asynq"
	"go.uber.org/zap/zapcore"
)

func main() {
	var env = bootstrap.Env{}
	bootstrap.NewEnv(&env)

	logConfig := pkglog.NewConfig()
	log := pkglog.InitLogger(logConfig, zapcore.DebugLevel, env.IsProduction())

	cf := asynq.Config{
		Concurrency: env.QUEUE.Concurrency,
		Queues:      env.QUEUE.Queues,
	}

	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     env.QUEUE.Addr,
			DB:       env.QUEUE.DB,
			Password: env.QUEUE.Password,
			Network:  env.QUEUE.Network,
		},
		cf,
	)
	mux := asynq.NewServeMux()
	// Register tasks and handlers
	worker.NewEmailSystem(mux, log)

	if err := srv.Run(mux); err != nil {
		log.Fatal("Could not run server: " + err.Error())
	}
}
