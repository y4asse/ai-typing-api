package usecase

import (
	"ai-typing/model"
	"ai-typing/repository"
)

type ICreatedTextUsecase interface {
	CreateCreatedText(createdText model.CreatedText) (model.CreatedTextResponse, error)
	GetAllCreatedTexts() ([]model.CreatedText, error)
	FindByGameId(gameId string) ([]model.CreatedText, error)
}

type createdTextUsecase struct {
	createdTextRepository repository.ICreatedTextRepository
}

func NewCreatedTextUsecase(createdTextRepository repository.ICreatedTextRepository) ICreatedTextUsecase {
	return &createdTextUsecase{createdTextRepository}
}

func (createdTextUsecase *createdTextUsecase) GetAllCreatedTexts() ([]model.CreatedText, error) {
	var createdTexts []model.CreatedText
	if err := createdTextUsecase.createdTextRepository.GetAllCreatedTexts(&createdTexts); err != nil {
		return nil, err
	}
	return createdTexts, nil
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
	}
	return resCreatedText, nil
}

func (createdTextUsecase *createdTextUsecase) FindByGameId(gameId string) ([]model.CreatedText, error) {
	createdTexts := []model.CreatedText{}
	err := createdTextUsecase.createdTextRepository.FindByGameId(&createdTexts, gameId)
	if err != nil {
		return nil, err
	}
	return createdTexts, nil
}
