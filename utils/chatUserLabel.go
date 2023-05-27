package utils

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
)

const (
	askUserLabel = "If a user registers for %splease speculate about the user's characteristics and provide at least 5 user tags, just output the result without any beginning, end, or explanation."
)

func ChatUserLabel(token, urls string) (openai.ChatCompletionResponse, error) {
	client := openai.NewClient(token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf(askUserLabel, urls),
				},
			},
		},
	)

	return resp, err
}
