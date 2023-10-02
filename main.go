package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type OpenAiMessagePayload struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIChatPayload struct {
	Model    string                   `json:"model"`
	Messages [2]*OpenAiMessagePayload `json:"messages"`
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	gptKey := os.Getenv("OPENAI_API_KEY")
	fmt.Println(gptKey)

	gptMessage := OpenAiMessagePayload{
		Role:    "system",
		Content: "You are a helpful assistant.",
	}

	userMessage := OpenAiMessagePayload{
		Role:    "user",
		Content: "Hello!",
	}

	payload := OpenAIChatPayload{
		Model: "gpt-3.5-turbo",
		Messages: [2]*OpenAiMessagePayload{
			&gptMessage,
			&userMessage,
		},
	}

	data, err := json.Marshal(payload)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Authorization", "Bearer "+gptKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}
