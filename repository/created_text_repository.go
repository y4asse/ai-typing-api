package repository

import (
	"ai-typing/model"

	"gorm.io/gorm"
)

type ICreatedTextRepository interface {
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
