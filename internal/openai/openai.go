package openai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

var (
	// ErrOpenAINoAnswer is returned when OpenAI does return OK but doesn't
	// return any answer (is not expected to happen).
	ErrOpenAINoAnswer = fmt.Errorf("openai: no answer")
)

// OpenAI is the OpenAI API client that we will use to perform OpenAI requests.
type OpenAI struct {
	openAI *openai.Client
}

// New creates a new OpenAI client.
func New(apiKey string) *OpenAI {
	return &OpenAI{
		openAI: openai.NewClient(apiKey),
	}
}

// ChatCompletion performs a chat completion request to OpenAI.
//
// It does use the GPT-3 0.5 Turbo model.
// It does use the User role for the message.
func (ai *OpenAI) ChatCompletion(
	ctx context.Context, content string) (string, error) {
	resp, err := ai.openAI.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)
	if err != nil {
		return "", fmt.Errorf("ai.openAI.CreateChatCompletion, err: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", ErrOpenAINoAnswer
	}

	return resp.Choices[0].Message.Content, nil
}
