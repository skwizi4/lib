package Gpt_3

type GPT3 interface {
	Request(text string) (string, error)
}
