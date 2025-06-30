package main

import (
	"cms-server/bootstrap"
	"cms-server/infrastructure/api/router"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.DB
	log := app.Log
	cacheApp := app.Cache
	defer db.Close()
	defer app.QueneClient.Close()

	fiberApp := fiber.New(fiber.Config{
		AppName:       env.NAME_APP,
		CaseSensitive: true,
		Prefork:       true,
		StrictRouting: true,
	})

	fiberApp.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	fiberApp.Use(cache.New((cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Get("No-Cache") == "true"
		},
		Expiration:   10 * time.Minute,
		CacheControl: true,
		Storage:      cacheApp,
	})))

	// Registering the route
	router.InitRouter(fiberApp, db, log, app.QueneClient, env, cacheApp)

	if err := fiberApp.Listen(":" + env.PORT_APP); err != nil {
		log.Fatal("Error starting the server: " + err.Error())
	}
}
