package usecase

import (
	"ai-typing/model"
	"ai-typing/repository"
)

type IGameUsecase interface {
	CreateGame(game model.Game) (model.GameResponse, error)
	GetGameRanking(border int) ([]model.GameResponse, error)
	GetGameHistory(userId string) ([]model.GameResponse, error)
	GetAllGame() ([]model.GameResponse, error)
	GetLatestGames(offset int) ([]model.GameResponse, error)
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

func (gameUsecase *gameUsecase) CreateGame(game model.Game) (model.GameResponse, error) {
	if err := gameUsecase.gameRepository.CreateGame(&game); err != nil {
		return model.GameResponse{}, err
	}
	resGame := model.GameResponse{
		ID:            game.ID,
		Score:         game.Score,
		InputedThema:  game.InputedThema,
		CreatedAt:     game.CreatedAt,
		ModeId:        game.ModeId,
		TotalKeyCount: game.TotalKeyCount,
		TotalMissType: game.TotalMissType,
		TotalTime:     game.TotalTime,
	}
	return resGame, nil
}

func (gameUsecase *gameUsecase) GetGameRanking(border int) ([]model.GameResponse, error) {
	games := []model.Game{}
	if err := gameUsecase.gameRepository.GetGameRanking(&games, border); err != nil {
		return nil, err
	}
	resGames := []model.GameResponse{}
	for _, v := range games {
		game := model.GameResponse{
			ID:            v.ID,
			Score:         v.Score,
			InputedThema:  v.InputedThema,
			CreatedAt:     v.CreatedAt,
			ModeId:        v.ModeId,
			TotalKeyCount: v.TotalKeyCount,
			TotalMissType: v.TotalMissType,
			TotalTime:     v.TotalTime,
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
			ID:            v.ID,
			Score:         v.Score,
			InputedThema:  v.InputedThema,
			CreatedAt:     v.CreatedAt,
			ModeId:        v.ModeId,
			TotalKeyCount: v.TotalKeyCount,
			TotalMissType: v.TotalMissType,
			TotalTime:     v.TotalTime,
		}
		resGames = append(resGames, game)
	}
	return resGames, nil
}

func (gameUsecase *gameUsecase) GetAllGame() ([]model.GameResponse, error) {
	games := []model.Game{}
	err := gameUsecase.gameRepository.GetAllGame(&games)
	if err != nil {
		return nil, err
	}
	resGames := []model.GameResponse{}
	for _, v := range games {
		game := model.GameResponse{
			ID:            v.ID,
			Score:         v.Score,
			InputedThema:  v.InputedThema,
			CreatedAt:     v.CreatedAt,
			ModeId:        v.ModeId,
			TotalKeyCount: v.TotalKeyCount,
			TotalMissType: v.TotalMissType,
			TotalTime:     v.TotalTime,
		}
		resGames = append(resGames, game)
	}
	return resGames, nil
}

func (gameUsecase *gameUsecase) GetLatestGames(offset int) ([]model.GameResponse, error) {
	games := []model.Game{}
	err := gameUsecase.gameRepository.GetLatestGames(&games, offset)
	if err != nil {
		return nil, err
	}
	resGames := []model.GameResponse{}
	for _, v := range games {
		game := model.GameResponse{
			ID:            v.ID,
			Score:         v.Score,
			InputedThema:  v.InputedThema,
			CreatedAt:     v.CreatedAt,
			ModeId:        v.ModeId,
			TotalKeyCount: v.TotalKeyCount,
			TotalMissType: v.TotalMissType,
			TotalTime:     v.TotalTime,
		}
		resGames = append(resGames, game)
	}
	return resGames, nil
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
