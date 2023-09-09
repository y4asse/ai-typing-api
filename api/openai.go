package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"ai-typing/model"
	"ai-typing/utils"
)

type Body struct {
	Model    string     `json:"model"`
	Messages []Messages `json:"messages"`
}

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func CreateAiText(thema string, detail string, aiModel string) (string, error) {
	method := "POST"
	OPEN_AI_URL := "https://api.openai.com/v1/chat/completions"
	messages := []Messages{
		{Role: "system", Content: "あなたはタイピングゲーム用の文章を作成するアシスタントです.「ああああああ」のように長押しで入力できる文は作成しません."},
		{Role: "user", Content: `「` + thema + `」についての文を5つ考えて{1:文章, 2:文章, 3:文章, 4:文章, 5:文章}のjson形式で教えて.`},
	}
	if detail == "を連打する文章" {
		messages = []Messages{
			{Role: "user", Content: "「無駄」を連打する短文を5つ考えて{1:文章, 2:文章, 3:文章, 4:文章, 5:文章}のjson形式で教えて"},
			{Role: "assistant", Content: thema + `{1: "無駄無駄", 2: "無駄無駄無駄無駄", 3: "無駄無駄無駄無駄", 4: "無駄無駄無駄無駄", 5: "無駄無駄無駄無駄無駄無駄"}`},
			{Role: "user", Content: "「" + thema + "」を連打する短文を5つ考えて{1:文章, 2:文章, 3:文章, 4:文章, 5:文章}のjson形式で教えて"},
		}
	}
	if detail == "文章" {
		messages = []Messages{
			{Role: "user", Content: thema + "短文を5つ考えて{1:文章, 2:文章, 3:文章, 4:文章, 5:文章}のjson形式で教えて"},
		}

	}
	reqBody := Body{
		Model:    aiModel,
		Messages: messages,
	}
	jsonBody, _ := json.Marshal(reqBody)
	API_KEY := os.Getenv("API_KEY")
	if API_KEY == "" {
		fmt.Println("API_KEYを設定してください")
		return "", fmt.Errorf("API_KEYを設定してください")
	}

	req, err := http.NewRequest(method, OPEN_AI_URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("リクエストの作成に失敗しました:", err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+API_KEY)
	clientOpenai := &http.Client{}
	resp, err := clientOpenai.Do(req)
	if err != nil {
		fmt.Println("リクエストの送信に失敗しました", err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) //bodyの読みとり
	if err != nil {
		fmt.Println("bodyの読み取りに失敗しました:", err)
		return "", err
	}
	data := model.OpenaiResponse{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("デシリアライズに失敗しました", err)
		return "", err
	}

	if data.Choices == nil {
		fmt.Println("Open AIからのレスポンスに問題があります")
		fmt.Println(resp.Status)
		return "", fmt.Errorf("OpenAiからのレスポンスに問題があります")
	}
	message := data.Choices[0].Message.Content
	message = utils.TrimOtherChar(message)

	return message, nil
}

func Analyse(time string, typeKeyCount string, missTypeCount string, kpm string, missTypeKey string, score string, accuracy string) (string, error) {
	aiModel := "gpt-4"
	method := "POST"
	OPEN_AI_URL := "https://api.openai.com/v1/chat/completions"
	messages := []Messages{
		{
			Role:    "system",
			Content: "あなたはタイピングゲームの分析を行うアシスタントです",
		},
		{
			Role:    "user",
			Content: "これらはタイピングゲームの結果です。分析して改善点とアドバイスと具体的な練習法を300文字程度で教えて！入力時間" + time + "秒,入力キー数" + typeKeyCount + ", ミス入力数" + missTypeCount + ",KPM" + kpm + ",正確率" + accuracy + "%, 間違えた文字" + missTypeKey,
		},
	}
	reqBody := Body{
		Model:    aiModel,
		Messages: messages,
	}

	API_KEY := os.Getenv("API_KEY")
	if API_KEY == "" {
		fmt.Println("API_KEYを設定してください")
		return "", fmt.Errorf("API_KEYを設定してください")
	}
	jsonBody, _ := json.Marshal(reqBody)
	req, err := http.NewRequest(method, OPEN_AI_URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("リクエストの作成に失敗しました:", err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+API_KEY)
	clientOpenai := &http.Client{}
	resp, err := clientOpenai.Do(req)
	if err != nil {
		fmt.Println("リクエストの送信に失敗しました", err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) //bodyの読みとり
	if err != nil {
		fmt.Println("bodyの読み取りに失敗しました:", err)
		return "", err
	}
	data := model.OpenaiResponse{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("デシリアライズに失敗しました", err)
		return "", err
	}
	message := data.Choices[0].Message.Content
	return message, nil
}
