package repository

import (
	"ai-typing/model"
	"fmt"

	"gorm.io/gorm"
)

type IGameRepository interface {
	CreateGame(game *model.Game) error
	GetGameRanking(games *[]model.Game) error
	GetGameHistory(game *[]model.Game, userId string) error
	GetAllGame(games *[]model.Game) error
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

func (gameRepository *gameRepository) GetGameHistory(games *[]model.Game, userId string) error {
	//gameからuser_id = userIdのデータを取得
	fmt.Println(userId)
	if err := gameRepository.db.Where("user_id = ?", userId).Order("created_at desc").Find(games).Error; err != nil {
		return err
	}
	return nil
}

func (gameRepository *gameRepository) GetAllGame(games *[]model.Game) error {
	if err := gameRepository.db.Find(games).Error; err != nil {
		return err
	}
	return nil
}