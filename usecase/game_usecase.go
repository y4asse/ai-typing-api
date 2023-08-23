package usecase

import (
	"ai-typing/model"
	"ai-typing/repository"
)

type IGameUsecase interface {
	CreateGame(game model.Game) (model.Game, error)
	GetGameRanking(border int) ([]model.Game, error)
	GetGameHistory(userId string) ([]model.Game, error)
	GetAllGame() ([]model.Game, error)
	GetLatestGames(offset int) ([]model.Game, error)
	GetTotalGameCount() (int64, error)
	UpdateGameScore(score int, totalKeyCount int, totalTime int, TotalMissType int, gameId string) error
}

type gameUsecase struct {
	gameRepository        repository.IGameRepository
	createdTextRepository repository.ICreatedTextRepository
}

func NewGameUsecase(gameRepository repository.IGameRepository, createdTextRepository repository.ICreatedTextRepository) IGameUsecase {
	return &gameUsecase{gameRepository, createdTextRepository}
}

func (gameUsecase *gameUsecase) CreateGame(game model.Game) (model.Game, error) {
	if err := gameUsecase.gameRepository.CreateGame(&game); err != nil {
		return model.Game{}, err
	}
	return game, nil
}

func (gameUsecase *gameUsecase) GetGameRanking(border int) ([]model.Game, error) {
	games := []model.Game{}
	if err := gameUsecase.gameRepository.GetGameRanking(&games, border); err != nil {
		return nil, err
	}
	return games, nil
}

func (gameUsecase *gameUsecase) GetGameHistory(userId string) ([]model.Game, error) {
	games := []model.Game{}
	err := gameUsecase.gameRepository.GetGameHistory(&games, userId)
	if err != nil {
		return nil, err
	}
	return games, nil
}

func (gameUsecase *gameUsecase) GetAllGame() ([]model.Game, error) {
	games := []model.Game{}
	err := gameUsecase.gameRepository.GetAllGame(&games)
	if err != nil {
		return nil, err
	}
	return games, nil
}

func (gameUsecase *gameUsecase) GetLatestGames(offset int) ([]model.Game, error) {
	games := []model.Game{}
	err := gameUsecase.gameRepository.GetLatestGames(&games, offset)
	if err != nil {
		return nil, err
	}
	return games, nil
}

func (gameUsecase *gameUsecase) GetTotalGameCount() (int64, error) {
	count, err := gameUsecase.gameRepository.GetTotalGameCount()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (gameUsecase *gameUsecase) UpdateGameScore(score int, totalKeyCount int, totalTime int, TotalMissType int, gameId string) error {
	err := gameUsecase.gameRepository.UpdateGameScore(score, totalKeyCount, totalTime, TotalMissType, gameId)
	if err != nil {
		return err
	}
	return nil
}
