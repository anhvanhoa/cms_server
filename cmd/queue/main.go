package main

import (
	"cms-server/bootstrap"
	"cms-server/internal/repository"
	"cms-server/internal/worker"
	pkglog "cms-server/pkg/logger"
	"cms-server/pkg/mailtemplate"

	"github.com/hibiken/asynq"
	"go.uber.org/zap/zapcore"
)

func main() {
	var env = bootstrap.Env{}
	bootstrap.NewEnv(&env)
	logConfig := pkglog.NewConfig()
	log := pkglog.InitLogger(logConfig, zapcore.DebugLevel, env.IsProduction())
	db := bootstrap.NewPostgresDB(&env, []any{}, log)
	defer db.Close()
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
	mailtemplate := mailtemplate.NewMailTemplate()
	mailProvider, err := bootstrap.NewMailProvider()
	if err != nil {
		log.Fatal("Could not create mail provider: " + err.Error())
	}
	// Register tasks and handlers
	worker.NewEmailSystem(
		mux,
		log,
		mailtemplate,
		mailProvider,
		repository.NewMailTplRepository(db),
		repository.NewMailProviderRepository(db),
		repository.NewMailHistoryRepository(db),
		repository.NewStatusHistoryRepository(db),
	)

	if err := srv.Run(mux); err != nil {
		log.Fatal("Could not run server: " + err.Error())
	}
}
