package usecase

import (
	"ai-typing/model"
	"ai-typing/repository"
	"fmt"
)

type IUserUsecase interface {
	FindByUserId(userId string) (model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
}

type userUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) IUserUsecase {
	return &userUsecase{userRepository}
}

func (userUsecase *userUsecase) FindByUserId(userId string) (model.User, error) {
	user := model.User{}
	if err := userUsecase.userRepository.FindByUserId(userId, &user); err != nil {
		fmt.Println(err.Error())
		return model.User{}, err
	}
	return user, nil
}

func (userUsecase *userUsecase) Create(user *model.User) error {
	if err := userUsecase.userRepository.Create(user); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (userUsecase *userUsecase) Update(user *model.User) error {
	if err := userUsecase.userRepository.Update(user); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
