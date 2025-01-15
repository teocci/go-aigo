// Package webserver
// Created by RTT.
// Author: teocci@yandex.com on 2025-1월-14
package webserver

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/teocci/aigo/src/config"
)

func Start(cfg *config.ServerSetup) {
	app := fiber.New()

	// Register routes
	registerRoutes(app)

	// Start server
	port := cfg.Web.Port
	log.Printf("Listening on :%d", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}

func registerRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to the modular Fiber app!",
		})
	})
}
