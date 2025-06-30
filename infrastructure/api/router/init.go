package router

import (
	"cms-server/bootstrap"
	pkglog "cms-server/infrastructure/service/logger"
	"cms-server/internal/service/cache"
	"cms-server/internal/service/queue"

	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	db    *pg.DB
	app   *fiber.App
	log   pkglog.Logger
	qc    queue.QueueClient
	env   *bootstrap.Env
	cache cache.RedisConfigImpl
}

func InitRouter(
	app *fiber.App,
	db *pg.DB,
	log pkglog.Logger,
	qc queue.QueueClient,
	env *bootstrap.Env,
	cache cache.RedisConfigImpl,
) {
	router := &Router{
		db:    db,
		app:   app,
		log:   log,
		qc:    qc,
		env:   env,
		cache: cache,
	}
	router.initAuthRouter()
}
