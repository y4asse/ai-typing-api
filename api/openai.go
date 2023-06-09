package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"ai-typing/model"
)

func CreateAiText(thema string) (string, error) {
	method := "POST"
	OPEN_AI_URL := "https://api.openai.com/v1/chat/completions"
	payload := strings.NewReader(`{"model": "gpt-3.5-turbo", "messages": [{"role": "user", "content": "` + thema + `について5つの短文を考えて{1:文章, 2:文章...}のjson形式で教えて"}]}`)
	API_KEY := os.Getenv("API_KEY")
	if API_KEY == "" {
		fmt.Println("API_KEYを設定してください")
		return "", fmt.Errorf("API_KEYを設定してください")
	}

	req, err := http.NewRequest(method, OPEN_AI_URL, payload)
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
	body, err := ioutil.ReadAll(resp.Body) //bodyの読みとり
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
		return "", fmt.Errorf("OpenAiからのレスポンスに問題があります")
	}
	message := data.Choices[0].Message.Content
	message = strings.Replace(message, "{", "", 1)
	message = strings.Replace(message, "}", "", 1)
	message = strings.ReplaceAll(message, "『", "")
	message = strings.ReplaceAll(message, "』", "")
	message = strings.ReplaceAll(message, "「", "")
	message = strings.ReplaceAll(message, "」", "")
	message = strings.ReplaceAll(message, "\n", "")
	message = strings.ReplaceAll(message, `"`, "")
	message = strings.ReplaceAll(message, `]`, "")
	message = strings.ReplaceAll(message, `[`, "")
	message = strings.ReplaceAll(message, ` `, "")
	message = strings.ReplaceAll(message, `　`, "")
	message = strings.ReplaceAll(message, `、`, "")
	message = strings.ReplaceAll(message, `。`, "")
	message = strings.ReplaceAll(message, `(`, "")
	message = strings.ReplaceAll(message, `)`, "")
	message = strings.ReplaceAll(message, `）`, "")
	message = strings.ReplaceAll(message, `（`, "")
	message = strings.ReplaceAll(message, `》`, "")
	message = strings.ReplaceAll(message, `《`, "")
	message = strings.ReplaceAll(message, `一`, "いち")
	message = strings.ReplaceAll(message, `二`, "に")
	message = strings.ReplaceAll(message, `三`, "さん")
	message = strings.ReplaceAll(message, `四`, "よん")
	message = strings.ReplaceAll(message, `五`, "ご")
	message = strings.ReplaceAll(message, `六`, "ろく")
	message = strings.ReplaceAll(message, `七`, "なな")
	message = strings.ReplaceAll(message, `八`, "はち")
	message = strings.ReplaceAll(message, `九`, "きゅう")
	message = strings.ReplaceAll(message, `十`, "じゅう")
	message = strings.ReplaceAll(message, `×`, "")
	message = strings.ReplaceAll(message, `・`, "")
	message = strings.ReplaceAll(message, `”`, "")
	message = strings.ReplaceAll(message, `“`, "")
	message = strings.ReplaceAll(message, `’`, "")
	message = strings.ReplaceAll(message, `‘`, "")
	message = strings.ReplaceAll(message, `：`, "")
	message = strings.ReplaceAll(message, `；`, "")
	message = strings.ReplaceAll(message, `｝`, "")
	message = strings.ReplaceAll(message, `｛`, "")
	message = strings.ReplaceAll(message, `＜`, "")
	message = strings.ReplaceAll(message, `＞`, "")
	message = strings.ReplaceAll(message, `＿`, "")
	message = strings.ReplaceAll(message, `＃`, "")
	message = strings.ReplaceAll(message, `＄`, "")
	message = strings.ReplaceAll(message, `％`, "")
	message = strings.ReplaceAll(message, `＆`, "")
	message = strings.ReplaceAll(message, `＝`, "")
	message = strings.ReplaceAll(message, `～`, "")
	message = strings.ReplaceAll(message, `＾`, "")
	message = strings.ReplaceAll(message, `￥`, "")
	message = strings.ReplaceAll(message, `｜`, "")

	return message, nil
}
