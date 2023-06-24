package repository

import (
	"ai-typing/model"

	"gorm.io/gorm"
)

type IGameRepository interface {
	CreateGame(game *model.Game) error
	GetGameRanking(games *[]model.Game) error
}

type gameRepository struct {
	db *gorm.DB
}

// Dependency Injectionで依存関係を注入
func NewGameRepository(db *gorm.DB) IGameRepository {
	return &gameRepository{db}
}

func (gameRepository *gameRepository) CreateGame(game *model.Game) error {
	if err := gameRepository.db.Create(game).Error; err != nil {
		return err
	}
	return nil
}
func (gameRepository *gameRepository) GetGameRanking(games *[]model.Game) error {
	if err := gameRepository.db.Order("score desc").Limit(10).Find(games).Error; err != nil {
		return err
	}
	return nil
}
