package openai_next_api

import (
	"context"
	"log"
)

func Chat(secret_key string, questions []string) []string {
	client := NewClient(secret_key)
	req := ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []ChatCompletionMessage{
			{
				Role:    "user",
				Content: "",
			},
		},
	}
	answers := make([]string, 0, len(questions))

	for _, question := range questions {
		req.Messages = append(req.Messages, ChatCompletionMessage{
			Role:    "user",
			Content: question,
		})
		resp, err := client.CreateChatCompletion(context.Background(), req)
		if err != nil {
			log.Printf("ChatCompletion error: %v\n", err)
			continue
		}
		answers = append(answers, resp.Choices[0].Message.Content)
		req.Messages = append(req.Messages, resp.Choices[0].Message)
	}

	return answers
}
