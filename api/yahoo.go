package api

import (
	"ai-typing/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	AppID = "dj00aiZpPU03Mlp2em91ejFnWiZzPWNvbnN1bWVyc2VjcmV0Jng9NzY-"
	URL   = "https://jlp.yahooapis.jp/FuriganaService/V2/furigana"
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
	headers := http.Header{
		"Content-Type": []string{"application/json"},
		"User-Agent":   []string{"Yahoo AppID: " + AppID},
	}

	paramDic := RequestBody{
		ID:      "1234-1",
		JsonRPC: "2.0",
		Method:  "jlp.furiganaservice.furigana",
		Params: Params{
			Q:     query,
			Grade: 1,
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
	fmt.Println("translating to hiragana...")
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
	fmt.Println("successfully translate")
	return hiragana, nil
}