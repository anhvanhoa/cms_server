package router

import (
	"cms-server/bootstrap"
	"cms-server/domain/service/cache"
	"cms-server/domain/service/queue"
	pkglog "cms-server/infrastructure/service/logger"

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
	router.initTypeMailRouter()
}
