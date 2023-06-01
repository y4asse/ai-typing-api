package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"ai-typing/model"
)

type IOpenaiUsecase interface {
	GetAiText(thema string) (model.AiTextResponse, error)
}

type openaiUsecase struct{}

func NewOpenaiUsecase() IOpenaiUsecase {
	return &openaiUsecase{}
}

func (ou *openaiUsecase) GetAiText(thema string) (model.AiTextResponse, error) {
	method := "POST"
	url := "https://api.openai.com/v1/chat/completions"
	payload := strings.NewReader(`{"model": "gpt-3.5-turbo", "messages": [{"role": "user", "content": "` + thema + `"}]}`)
	API_KEY := os.Getenv("API_KEY")
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println("リクエストの作成に失敗しました:", err)
		return model.AiTextResponse{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+API_KEY)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("リクエストの送信に失敗しました", err)
		return model.AiTextResponse{}, err
	}
	defer resp.Body.Close()
	fmt.Println("resp:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("bodyの読み取りに失敗しました:", err)
		return model.AiTextResponse{}, err
	}
	data := model.OpenaiResponse{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("デシリアライズに失敗しました", err)
	}
	message := data.Choices[0].Message.Content
	fmt.Println("message:", message)
	//TODOmessageを分割する

	resAiText := model.AiTextResponse{
		Text: message,
	}
	return resAiText, nil
}
