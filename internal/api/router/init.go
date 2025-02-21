package router

import (
	pkglog "cms-server/pkg/logger"

	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	db  *pg.DB
	app *fiber.App
	log pkglog.Logger
}

func InitRouter(app *fiber.App, db *pg.DB, log pkglog.Logger) {
	router := &Router{
		db:  db,
		app: app,
		log: log,
	}
	router.initAuthRouter()
}
