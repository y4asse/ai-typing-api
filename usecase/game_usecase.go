package usecase

import (
	"ai-typing/model"
	"ai-typing/repository"
)

type IGameUsecase interface {
	CreateGame(game model.Game) (model.GameResponse, error)
	GetGameRanking() ([]model.GameResponse, error)
	GetGameHistory(userId string) ([]model.GameResponse, error)
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

func (gameUsecase *gameUsecase) GetGameRanking() ([]model.GameResponse, error) {
	games := []model.Game{}
	if err := gameUsecase.gameRepository.GetGameRanking(&games); err != nil {
		return nil, err
	}
	resGames := []model.GameResponse{}
	for _, v := range games {
		game := model.GameResponse{
			ID:           v.ID,
			UserId:       v.UserId,
			Score:        v.Score,
			InputedThema: v.InputedThema,
			CreatedAt:    v.CreatedAt,
			ModeId:       v.ModeId,
		}
		resGames = append(resGames, game)
	}
	return resGames, nil
}

func (gameUsecase *gameUsecase) GetGameHistory(userId string) ([]model.GameResponse, error) {
	games := []model.Game{}
	err := gameUsecase.gameRepository.GetGameHistory(&games, userId)
	if err != nil {
		return nil, err
	}
	resGames := []model.GameResponse{}
	for _, v := range games {
		game := model.GameResponse{
			ID:           v.ID,
			UserId:       v.UserId,
			Score:        v.Score,
			InputedThema: v.InputedThema,
			CreatedAt:    v.CreatedAt,
			ModeId:       v.ModeId,
		}
		resGames = append(resGames, game)
	}
	return resGames, nil
}
