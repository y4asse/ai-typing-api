package usecase

import (
	"ai-typing/model"
	"ai-typing/repository"
)

type ILikeUsecase interface {
	FetchAll() ([]model.Like, error)
	Create(like model.Like) error
	Delete(gameId string) error
	FetchAllByGameId(gameId string) ([]model.Like, error)
	GetNumByGameId(gameId string) (int, error)
}
type likeUsecase struct {
	likeRepository repository.IlikeRepository
}

func NewLikeUsecase(likeRepository repository.IlikeRepository) ILikeUsecase {
	return &likeUsecase{likeRepository}
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
