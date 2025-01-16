// Package webserver
// Created by RTT.
// Author: teocci@yandex.com on 2025-1월-15
package webserver

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/teocci/aigo/src/models"
)

func registerRoutes(app *fiber.App) {
	// API endpoint to capture and return the JSON payload
	app.Post("/api/v1/completion", handleCompletion)
}

func handleCompletion(c *fiber.Ctx) error {
	// Parse the incoming request payload
	var requestBody models.CompletionRequest
	if err := c.BodyParser(&requestBody); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON payload",
		})
	}

	append(requestBody.Messages, models.ModePrompts[])

	// Log the incoming request
	log.Printf("Received request: %+v", requestBody)

	// Call OpenAI API using the service layer
	response, err := models.GetXAICompletion(requestBody)
	if err != nil {
		log.Printf("Error calling OpenAI API: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process the request",
		})
	}

	// Stream or return the response
	return c.JSON(response)
}
