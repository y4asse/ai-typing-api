package usecase

import (
	"ai-typing/api"
	"ai-typing/model"
	"fmt"
	"strings"
	"unicode/utf8"
)

type IOpenaiUsecase interface {
	GetAiText(thema string, detail string, aiModel string) (model.AiTextResponse, error)
	Analyse(model.AnalyseRequest) (string, error)
}

type openaiUsecase struct{}

func NewOpenaiUsecase() IOpenaiUsecase {
	return &openaiUsecase{}
}

func (ou *openaiUsecase) GetAiText(thema string, detail string, aiModel string) (model.AiTextResponse, error) {
	trimThema := strings.ReplaceAll(thema, " ", "")
	if utf8.RuneCountInString(trimThema) > 10 {
		fmt.Println(thema, utf8.RuneCountInString(trimThema))
		return model.AiTextResponse{}, fmt.Errorf("テーマは10文字以内で入力してください")
	}
	text, err := api.CreateAiText(thema, detail, aiModel)
	if err != nil {
		fmt.Println("aiテキストの作成に失敗しました", err)
		return model.AiTextResponse{}, err
	}
	hiragana, err := api.Post(text)
	if err != nil {
		fmt.Println("ひらがなへの変換に失敗しました", err)
		return model.AiTextResponse{}, err
	}
	hiraganaArr := strings.Split(hiragana, ",")
	textArr := strings.Split(text, ",")
	for i, text := range textArr {
		if !strings.Contains(text, ":") {
			fmt.Println("AIが作成したテキストに問題があります")
			fmt.Println(textArr)
			return model.AiTextResponse{}, fmt.Errorf("AIが作成したテキストに問題があります")
		}
		textArr[i] = strings.Split(text, ":")[1]
	}
	for i, hiragana := range hiraganaArr {
		if !strings.Contains(hiragana, ":") {
			fmt.Println("AIが作成したテキストに問題があります(hiragana)")
			return model.AiTextResponse{}, fmt.Errorf("AIが作成したテキストに問題があります(hiragana)")
		}
		hiraganaArr[i] = strings.Split(hiragana, ":")[1]
	}

	resAiTextResponse := model.AiTextResponse{
		Text:     textArr,
		Hiragana: hiraganaArr,
	}
	return resAiTextResponse, nil
}

func (ou *openaiUsecase) Analyse(requestBody model.AnalyseRequest) (string, error) {
	time := requestBody.Time
	typeKeyCount := requestBody.TypeKeyCount
	missTypeCount := requestBody.MissTypeCount
	kpm := requestBody.KPM
	missTypeKey := requestBody.MissTypeKey
	score := requestBody.Score
	accuracy := requestBody.Accuracy
	analyseRes, err := api.Analyse(time, typeKeyCount, missTypeCount, kpm, missTypeKey, score, accuracy)
	if err != nil {
		fmt.Println("解析に失敗しました", err)
		return "", err
	}
	return analyseRes, nil
}
