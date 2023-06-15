package usecase

import (
	"ai-typing/api"
	"ai-typing/model"
	"fmt"
	"strings"
)

type IOpenaiUsecase interface {
	GetAiText(thema string) (model.AiTextResponse, error)
}

type openaiUsecase struct{}

func NewOpenaiUsecase() IOpenaiUsecase {
	return &openaiUsecase{}
}

func (ou *openaiUsecase) GetAiText(thema string) (model.AiTextResponse, error) {

	text, err := api.CreateAiText(thema)
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
