// Package webserver
// Created by RTT.
// Author: teocci@yandex.com on 2025-1월-15
package webserver

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func registerRoutes(app *fiber.App) {
	// API endpoint to capture and return the JSON payload
	app.Post("/api/search", func(c *fiber.Ctx) error {
		// Read and log the incoming JSON
		var requestBody map[string]interface{}
		if err := c.BodyParser(&requestBody); err != nil {
			log.Printf("Error parsing request body: %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON payload",
			})
		}

		log.Printf("Received request body: %v", requestBody)

		// Stream the response back
		return c.JSON(requestBody)
	})
}
