package controller

import (
	"ai-typing/model"
	"ai-typing/usecase"
	"fmt"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
)

type IUserController interface {
	GetUser(context echo.Context) error
	Update(context echo.Context) error
}

type userController struct {
	userUsecase usecase.IUserUsecase
}

func NewUserController(userUsecase usecase.IUserUsecase) IUserController {
	return &userController{userUsecase}
}

func (userController *userController) GetUser(context echo.Context) error {
	token := context.Get("token").(*auth.Token)
	claims := token.Claims
	uid, _ := claims["user_id"].(string)
	user, err := userController.userUsecase.FindByUserId(uid)
	name, _ := claims["name"].(string)
	image, _ := claims["picture"].(string)
	userModel := model.User{
		UserId: uid,
		Name:   name,
		Image:  image,
	}
	if err != nil {
		if err.Error() == "record not found" {
			//データが無かったら新規作成
			err := userController.userUsecase.Create(&userModel)
			if err != nil {
				fmt.Println(err.Error())
				return context.JSON(http.StatusInternalServerError, err.Error())
			}
			return context.JSON(http.StatusOK, userModel)
		}
		fmt.Println(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, user)
}

func (userController *userController) Update(context echo.Context) error {
	token := context.Get("token").(*auth.Token)
	claims := token.Claims
	uid, _ := claims["user_id"].(string)
	newUser := model.User{
		UserId: uid,
	}
	if err := context.Bind(&newUser); err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	if err := userController.userUsecase.Update(&newUser); err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, newUser)
}
