package Gpt_3

import "github.com/go-resty/resty/v2"

type (
	GptClient struct {
		model  string
		APIKey string
		client *resty.Client
	}
	ChatResponse struct {
		ID      string   `json:"id"`
		Choices []choice `json:"choices"`
	}
	choice struct {
		Message message `json:"message"`
	}
	message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}
)
