package usecase

import (
	"ai-typing/model"
	"ai-typing/repository"
)

type ICreatedTextUsecase interface {
	CreateCreatedText(createdText model.CreatedText) (model.CreatedTextResponse, error)
}

type createdTextUsecase struct {
	createdTextRepository repository.ICreatedTextRepository
}

func NewCreatedTextUsecase(createdTextRepository repository.ICreatedTextRepository) ICreatedTextUsecase {
	return &createdTextUsecase{createdTextRepository}
}

func (createdTextUsecase *createdTextUsecase) CreateCreatedText(createdText model.CreatedText) (model.CreatedTextResponse, error) {
	if err := createdTextUsecase.createdTextRepository.CreateCreatedText(&createdText); err != nil {
		return model.CreatedTextResponse{}, err
	}
	resCreatedText := model.CreatedTextResponse{
		ID:       createdText.ID,
		Text:     createdText.Text,
		Hiragana: createdText.Hiragana,
		GameId:   createdText.GameId,
		IsPost:   createdText.IsPost,
	}
	return resCreatedText, nil
}
