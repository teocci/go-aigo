// Package models
// Created by RTT.
// Author: teocci@yandex.com on 2025-1월-17
package models

import (
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type CompletionRequest struct {
	Messages []struct {
		Messages []openai.ChatCompletionMessage `json:"messages"`
	}
	Model string `json:"model"`
	Mode  string `json:"mode"`
}

// AddSystemPrompt enriches the request with the system prompt
func (req *CompletionRequest) AddSystemPrompt() ([]openai.ChatCompletionMessage, error) {
	// Fetch the system prompt for the specified group
	systemPrompt, exists := ModePrompts[req.Mode]
	if !exists {
		return nil, fmt.Errorf("no system prompt found for group: %s", req.Mode)
	}

	// Create the system prompt message
	systemMessage := openai.ChatCompletionMessage{
		Role:    "system",
		Content: systemPrompt,
	}

	// Prepend the system message to the user messages
	messages := append([]openai.ChatCompletionMessage{systemMessage}, req.Messages...)
	return messages, nil
}
