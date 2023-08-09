package usecase

import (
	"ai-typing/model"
	"ai-typing/repository"
	"fmt"
)

type ILikeUsecase interface {
	FetchAll() ([]model.Like, error)
	Create(like model.Like) error
	Delete(gameId string) error
	FetchAllByGameId(gameId string) ([]model.Like, error)
	GetNumByGameId(gameId string) (int, error)
	GetCountGroupByGameIdOrder(offset int, limit int) ([]model.GameWithCount, error)
}
type likeUsecase struct {
	likeRepository repository.IlikeRepository
	gameRepository repository.IGameRepository
}

func NewLikeUsecase(likeRepository repository.IlikeRepository, gameRepository repository.IGameRepository) ILikeUsecase {
	return &likeUsecase{likeRepository, gameRepository}
}

func (likeUsecase *likeUsecase) FetchAll() ([]model.Like, error) {
	likes := []model.Like{}
	err := likeUsecase.likeRepository.FetchAll(&likes)
	if err != nil {
		return nil, err
	}
	return likes, nil
}

func (likeUsecase *likeUsecase) Create(like model.Like) error {
	err := likeUsecase.likeRepository.Create(&like)
	if err != nil {
		return err
	}
	return nil
}

func (likeUsecase *likeUsecase) Delete(gameId string) error {
	err := likeUsecase.likeRepository.Delete(gameId)
	if err != nil {
		return err
	}
	return nil
}

func (likeUsecase *likeUsecase) FetchAllByGameId(gameId string) ([]model.Like, error) {
	likes := []model.Like{}
	err := likeUsecase.likeRepository.FetchAllByGameId(gameId, &likes)
	if err != nil {
		return nil, err
	}
	return likes, nil
}

func (likeUsecase *likeUsecase) GetNumByGameId(gameId string) (int, error) {
	num, err := likeUsecase.likeRepository.GetNumByGameId(gameId)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (likeUsecase *likeUsecase) GetCountGroupByGameIdOrder(offset int, limit int) ([]model.GameWithCount, error) {
	//gameIdを取得
	gameIdCounts := []model.GameIdCount{}
	err := likeUsecase.likeRepository.GetCountGroupByGameIdOrder(offset, limit, &gameIdCounts)
	if err != nil {
		return nil, err
	}

	//gameIdを元にgameを取得
	gameWithCounts := []model.GameWithCount{}
	for _, gameIdCount := range gameIdCounts {
		fmt.Println(gameIdCount.GameId)
		fmt.Println(gameIdCount.Count)
		game := model.Game{}
		//gameRepositoryfindを作る
		err := likeUsecase.gameRepository.FindOne(&game, gameIdCount.GameId)
		if err != nil {
			return nil, err
		}
		gameWithCount := model.GameWithCount{
			Game:  game,
			Count: gameIdCount.Count,
		}
		gameWithCounts = append(gameWithCounts, gameWithCount)
	}

	return gameWithCounts, nil
}
