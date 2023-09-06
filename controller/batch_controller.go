package controller

import (
	"ai-typing/usecase"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
)

type IBatchController interface {
	GetAllByUserId(context echo.Context) error
}

type batchController struct {
	batchUsecase usecase.IBatchUsecase
}

func NewBatchController(batchUsecase usecase.IBatchUsecase) IBatchController {
	return &batchController{batchUsecase}
}

func (batchController *batchController) GetAllByUserId(context echo.Context) error {
	token := context.Get("token").(*auth.Token)
	claims := token.Claims
	uid, _ := claims["user_id"].(string)
	batches, err := batchController.batchUsecase.GetAllByUserId(uid)
	if err != nil {
		return context.JSON(500, err.Error())
	}
	return context.JSON(200, batches)
}
