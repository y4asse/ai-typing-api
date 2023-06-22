package usecase

import (
	"ai-typing/model"
	"ai-typing/repository"
)

type IGameUsecase interface {
	CreateGame(game model.Game) (model.GameResponse, error)
}

type gameUsecase struct {
	gameRepository repository.IGameRepository
}

func NewGameUsecase(gameRepository repository.IGameRepository) IGameUsecase {
	return &gameUsecase{gameRepository}
}

func (gameUsecase *gameUsecase) CreateGame(game model.Game) (model.GameResponse, error) {
	if err := gameUsecase.gameRepository.CreateGame(&game); err != nil {
		return model.GameResponse{}, err
	}
	resGame := model.GameResponse{
		ID:           game.ID,
		UserId:       game.UserId,
		Score:        game.Score,
		InputedThema: game.InputedThema,
		CreatedAt:    game.CreatedAt,
		ModeId:       game.ModeId,
	}
	return resGame, nil
}
