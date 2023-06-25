package api

import (
	"ai-typing/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	URL = "https://jlp.yahooapis.jp/FuriganaService/V2/furigana"
)

type Params struct {
	Q     string `json:"q"`
	Grade int    `json:"grade"`
}

type RequestBody struct {
	ID      string `json:"id"`
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
}

func Post(query string) (string, error) {
	AppID := os.Getenv("CLIENT_ID")
	if AppID == "" {
		fmt.Println("CLIENT_IDを設定してください")
		return "", fmt.Errorf("CLIENT_IDを設定してください")
	}
	headers := http.Header{
		"Content-Type": []string{"application/json"},
		"User-Agent":   []string{"Yahoo AppID: " + AppID},
	}

	paramDic := RequestBody{
		ID:      "1234-1",
		JsonRPC: "2.0",
		Method:  "jlp.furiganaservice.furigana",
		Params: Params{
			Q: query,
		},
	}

	params, err := json.Marshal(paramDic)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(params))
	if err != nil {
		return "", err
	}
	req.Header = headers

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	data := model.YahooRes{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("デシリアライズに失敗しました", err)
		return "", err
	}

	var hiragana string
	for _, word := range data.Result.Word {
		if word.Furigana == "" {
			hiragana += word.Surface
		} else {
			hiragana += word.Furigana
		}
	}
	return hiragana, nil
}
