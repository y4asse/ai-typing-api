package repository

import (
	"ai-typing/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IBatchRepository interface {
	Create(batch *[]model.Batch) error
	GetAllByUserId(batchs *[]model.Batch, userId string) error
}

type batchRepository struct {
	db *gorm.DB
}

func NewBatchRepository(db *gorm.DB) IBatchRepository {
	return &batchRepository{db}
}

func (batchRepository *batchRepository) Create(batch *[]model.Batch) error {
	//重複時はエラーを無視
	result := batchRepository.db.Clauses(clause.OnConflict{DoNothing: true}).Create(batch)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}

func (batchRepository *batchRepository) GetAllByUserId(batchs *[]model.Batch, userId string) error {
	if err := batchRepository.db.Where("user_id = ?", userId).Find(batchs).Error; err != nil {
		return err
	}
	return nil
}
