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
	trimedHiragana := strings.ReplaceAll(hiragana, "一", "いち")
	trimedHiragana = strings.ReplaceAll(trimedHiragana, "数千", "すうせん")
	trimedHiragana = strings.ReplaceAll(trimedHiragana, "数万", "すうまん")
	trimedHiragana = strings.ReplaceAll(trimedHiragana, "数十万", "すうじゅうまん")
	trimedHiragana = strings.ReplaceAll(trimedHiragana, "二", "に")
	trimedHiragana = strings.ReplaceAll(trimedHiragana, "三", "さん")
	trimedHiragana = strings.ReplaceAll(trimedHiragana, "四", "よん")
	trimedHiragana = strings.ReplaceAll(trimedHiragana, "五", "ご")
	trimedHiragana = strings.ReplaceAll(trimedHiragana, "六", "ろく")
	trimedHiragana = strings.ReplaceAll(trimedHiragana, "七", "なな")
	trimedHiragana = strings.ReplaceAll(trimedHiragana, "八", "はち")
	trimedHiragana = strings.ReplaceAll(trimedHiragana, "九", "きゅう")
	trimedHiragana = strings.ReplaceAll(trimedHiragana, "十", "じゅう")
	trimedHiragana = strings.ReplaceAll(trimedHiragana, "百", "ひゃく")
	trimedHiragana = strings.ReplaceAll(trimedHiragana, "千", "せん")
	trimedHiragana = strings.ReplaceAll(trimedHiragana, "数", "かず")
	hiraganaArr := strings.Split(trimedHiragana, ",")
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
