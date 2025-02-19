package ai

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	openai "github.com/sashabaranov/go-openai"
)

func UseChatGPT4(s *discordgo.Session, m *discordgo.MessageCreate, cfg string) {

	if m.Content != "!openAI" {
		return
	}

	client := openai.NewClient(cfg)
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
	s.ChannelMessageSend("1015477578959699978", resp.Choices[0].Message.Content)
}
