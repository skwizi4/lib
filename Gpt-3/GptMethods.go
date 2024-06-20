package Gpt_3

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func InitGP3(model, APIKey string) GPT3 {
	return GptClient{APIKey: APIKey, model: model, client: resty.New()}
}
func (g GptClient) Request(text string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"
	data := map[string]interface{}{
		"model": g.model,
		"messages": []map[string]string{
			{"role": "user", "content": text},
		},
	}
	resp, err := g.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", g.APIKey)).
		SetBody(data).
		Post(url)

	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}

	// Проверка статуса ответа
	if resp.StatusCode() != 200 {
		return "", fmt.Errorf("request failed with status code %d: %s", resp.StatusCode(), resp.String())
	}

	var result ChatResponse
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return "", fmt.Errorf("error decoding response JSON: %v", err)
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("error: no choices in response")
	}

	return result.Choices[0].Message.Content, nil
}
