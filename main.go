package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	tok := os.Getenv("OPENAI_API_KEY")
	if tok == "" {
		panic("OPENAI_API_KEY is not set")
	}

	client := openai.NewClient(tok)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
