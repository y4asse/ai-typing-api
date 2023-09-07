package usecase

import (
	"ai-typing/api"
	"ai-typing/model"
	"ai-typing/utils"
	"ai-typing/validator"
	"fmt"
	"strings"
)

type IOpenaiUsecase interface {
	GetAiText(model.ThemaRequest) (model.AiTextResponse, error)
	Analyse(model.AnalyseRequest) (string, error)
}

type openaiUsecase struct {
	aiTextValidator validator.IAiTextValidator
}

func NewOpenaiUsecase(aiTextValidator validator.IAiTextValidator) IOpenaiUsecase {
	return &openaiUsecase{aiTextValidator}
}

func (ou *openaiUsecase) GetAiText(themaRequest model.ThemaRequest) (model.AiTextResponse, error) {
	themaRequest.Thema = strings.ReplaceAll(themaRequest.Thema, " ", "")
	//validation
	if err := ou.aiTextValidator.ThemaValidator(themaRequest); err != nil {
		return model.AiTextResponse{}, err
	}
	thema := themaRequest.Thema
	detail := themaRequest.Detail
	aiModel := themaRequest.AiModel

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
	trimedHiragana := utils.TrimNumKanji(hiragana)
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
