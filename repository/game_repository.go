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
	GetCreatedText(text *[]model.CreatedText, gameId string) error
	GetLatestGames(games *[]model.Game, offset int) error
	GetTotalGameCount() (int64, error)
	UpdateGameScore(score int, gameId string) error
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

func (gameRepository *gameRepository) GetCreatedText(text *[]model.CreatedText, gameId string) error {
	if err := gameRepository.db.Where("game_id = ?", gameId).Find(text).Error; err != nil {
		return err
	}
	return nil
}

func (gameRepository *gameRepository) GetLatestGames(games *[]model.Game, offset int) error {
	if err := gameRepository.db.Order("created_at desc").Offset(offset).Limit(10).Find(games).Error; err != nil {
		return err
	}
	return nil
}

func (gameRepository *gameRepository) GetTotalGameCount() (int64, error) {
	var totalGameCount int64
	if err := gameRepository.db.Model(&model.Game{}).Count(&totalGameCount).Error; err != nil {
		return 0, err
	}
	return totalGameCount, nil
}

func (gameRepository *gameRepository) UpdateGameScore(score int, gameId string) error {
	result := gameRepository.db.Model(&model.Game{}).Where("id = ?", gameId).Update("score", score)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
