package repository

import (
	"ai-typing/model"

	"gorm.io/gorm"
)

type IlikeRepository interface {
	FetchAll(likes *[]model.Like) error
	Create(like *model.Like) error
	Delete(gameId string) error
	FetchAllByGameId(gameId string, like *[]model.Like) error
	GetNumByGameId(gameId string) (int, error)
}

type likeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) IlikeRepository {
	return &likeRepository{db}
}

func (likeRepository *likeRepository) FetchAll(likes *[]model.Like) error {
	if err := likeRepository.db.Find(likes).Error; err != nil {
		return err
	}
	return nil
}

func (likeRepository *likeRepository) Create(like *model.Like) error {
	if err := likeRepository.db.Create(like).Error; err != nil {
		return err
	}
	return nil
}

func (likeRepository *likeRepository) Delete(gameId string) error {
	if err := likeRepository.db.Where("game_id = ?", gameId).Delete(&model.Like{}).Error; err != nil {
		return err
	}
	return nil
}

func (likeRepository *likeRepository) FetchAllByGameId(gameId string, likes *[]model.Like) error {
	if err := likeRepository.db.Where("game_id = ?", gameId).Find(likes).Error; err != nil {
		return err
	}
	return nil
}

func (likeRepository *likeRepository) GetNumByGameId(gameId string) (int, error) {
	var num int64
	if err := likeRepository.db.Model(&model.Like{}).Where("game_id = ?", gameId).Count(&num).Error; err != nil {
		return 0, err
	}
	return int(num), nil
}
