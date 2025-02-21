package bootstrap

import (
	"cms-server/internal/entity"
	pkglog "cms-server/pkg/logger"

	"github.com/go-pg/pg/v10"
	"go.uber.org/zap/zapcore"
)

type Application struct {
	Env *Env
	DB  *pg.DB
	Log pkglog.Logger
}

func App() *Application {
	env := NewEnv()

	logConfig := pkglog.NewConfig()
	log := pkglog.InitLogger(logConfig, zapcore.DebugLevel, env.IsProduction())

	entities := []interface{}{
		new(entity.User),
	}

	db := NewPostgresDB(env, entities)
	RegisterValidator()
	return &Application{
		Env: env,
		DB:  db,
		Log: log,
	}
}

func (app *Application) ClosePostgresDB() {
	ClosePostgresDB(app.DB)
}
