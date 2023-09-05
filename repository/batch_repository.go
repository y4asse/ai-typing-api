package repository

import (
	"ai-typing/model"

	"gorm.io/gorm"
)

type IBatchRepository interface {
	Create(batch *model.Batch) error
}

type batchRepository struct {
	db *gorm.DB
}

func NewBatchRepository(db *gorm.DB) IBatchRepository {
	return &batchRepository{db}
}

func (batchRepository *batchRepository) Create(batch *model.Batch) error {
	if err := batchRepository.db.Create(batch).Error; err != nil {
		return err
	}
	return nil
}
