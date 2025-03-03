package main

import (
	"cms-server/bootstrap"
	"cms-server/internal/api/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.DB
	log := app.Log
	tm := app.TM

	defer db.Close()
	defer app.QueneClient.Close()

	fiberApp := fiber.New(fiber.Config{
		AppName:       env.NAME_APP,
		CaseSensitive: true,
		Prefork:       true,
		StrictRouting: true,
	})

	// Registering the route
	router.InitRouter(fiberApp, db, log, app.QueneClient, tm)

	if err := fiberApp.Listen(":" + env.PORT_APP); err != nil {
		log.Fatal("Error starting the server: " + err.Error())
	}
}
