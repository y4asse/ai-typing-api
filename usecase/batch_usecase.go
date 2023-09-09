package usecase

import (
	"ai-typing/model"
	"ai-typing/repository"
)

type IBatchUsecase interface {
	GetAllByUserId(userId string) ([]model.Batch, error)
}

type batchUsecase struct {
	batchRepository repository.IBatchRepository
}

func NewBatchUsecase(batchRepository repository.IBatchRepository) IBatchUsecase {
	return &batchUsecase{batchRepository}
}

func (batchUsecase *batchUsecase) GetAllByUserId(userId string) ([]model.Batch, error) {
	batchs := []model.Batch{}
	if err := batchUsecase.batchRepository.GetAllByUserId(&batchs, userId); err != nil {
		return nil, err
	}
	return batchs, nil
}
