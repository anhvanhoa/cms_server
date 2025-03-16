package router

import (
	"cms-server/bootstrap"
	pkglog "cms-server/pkg/logger"

	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	db  *pg.DB
	app *fiber.App
	log pkglog.Logger
	qr  bootstrap.QueueClient
	env *bootstrap.Env
}

func InitRouter(
	app *fiber.App,
	db *pg.DB,
	log pkglog.Logger,
	qr bootstrap.QueueClient,
	env *bootstrap.Env,
) {
	router := &Router{
		db:  db,
		app: app,
		log: log,
		qr:  qr,
		env: env,
	}
	router.initAuthRouter()
}
