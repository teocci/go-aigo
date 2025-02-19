// Package models
// Created by RTT.
// Author: teocci@yandex.com on 2025-1월-16
package models

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"

	"github.com/teocci/aigo/src/env"
)

const (
	xaiBaseURL   = "https://api.x.ai/v1"
	defaultModel = "grok-2-1212"
)

var config openai.ClientConfig

func GetXAICompletion(req CompletionRequest) (interface{}, error) {
	err := env.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	xaiAPIKey := os.Getenv("XAI_API_KEY")

	fmt.Printf("req: %v \n", req)

	config = openai.DefaultConfig(xaiAPIKey)
	config.BaseURL = xaiBaseURL

	// Initialize OpenAI client
	client := openai.NewClientWithConfig(config)

	// Construct OpenAI request
	messages := make([]openai.ChatCompletionMessage, len(req.Messages))
	for i, msg := range req.Messages {
		messages[i] = openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}

	openAIReq := openai.ChatCompletionRequest{
		Model:    req.Model,
		Messages: messages,
		Stream:   true,
	}

	// Send request to OpenAI
	stream, err := client.CreateChatCompletionStream(context.Background(), openAIReq)
	if err != nil {
		log.Printf("OpenAI API error: %v", err)
		return nil, err
	}
	defer stream.Close()

	// Process the streamed response
	var result []string
	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("Stream finished")
			break
		}

		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			break
		}

		fmt.Printf("Stream response: %v\n", resp)

		result = append(result, resp.Choices[0].Delta.Content)
	}

	// Return the accumulated response
	return map[string]interface{}{
		"response": result,
	}, nil
}
