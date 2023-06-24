package repository

import (
	"ai-typing/model"

	"gorm.io/gorm"
)

type ICreatedTextRepository interface {
	// GetCreatedTextRepositoryRanking() error
	CreateCreatedText(createdText *model.CreatedText) error
}

type createdTextRepository struct {
	db *gorm.DB
}

// Dependency Injectionで依存関係を注入
func NewCreatedTextRepository(db *gorm.DB) ICreatedTextRepository {
	return &createdTextRepository{db}
}

func (createdTextRepository *createdTextRepository) CreateCreatedText(createdText *model.CreatedText) error {
	if err := createdTextRepository.db.Create(createdText).Error; err != nil {
		return err
	}
	return nil
}

// func (createdTextRepository *createdTextRepository) GetCreatedTextRepositoryRanking() error {
// 	if err := createdTextRepository.db.Order("score").Limit(10).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
