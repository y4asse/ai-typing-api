package repository

import (
	"ai-typing/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindByUserId(userId string, user *model.User) error
	Create(user *model.User) error
	Update(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (userRepository *userRepository) FindByUserId(userId string, user *model.User) error {
	if err := userRepository.db.Where("user_id = ?", userId).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (userRepository *userRepository) Create(user *model.User) error {
	if err := userRepository.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (userRepository *userRepository) Update(user *model.User) error {
	//Saveは更新するカラムを指定しないと全てのカラムを更新するので注意．
	//たとえばcreated_atを指定しないと0になってしまうのでSaveではなくUpdatesを使うことで空のカラムは更新されない
	// if err := userRepository.db.Model(&user).Where("user_id = ?", user.UserId).Save(&user).Error; err != nil {
	// 	return err
	// }
	if err := userRepository.db.Model(&user).Where("user_id = ?", user.UserId).Updates(&user).Error; err != nil {
		return err
	}
	if err := userRepository.db.First(&user, "user_id = ?", user.UserId).Error; err != nil {
		return err
	}
	return nil
}
