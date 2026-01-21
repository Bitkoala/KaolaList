package ai

import (
	"context"
)

// Request defines common AI request parameters
type Request struct {
	Prompt string
	Model  string
	Image  []byte // For vision tasks
}

// Response defines common AI response structure
type Response struct {
	Content string
	Error   error
}

// Provider defines the interface for AI models
type Provider interface {
	Chat(ctx context.Context, req Request) Response
	Translate(ctx context.Context, text string, targetLang string) Response
	Summarize(ctx context.Context, content string) Response
}
