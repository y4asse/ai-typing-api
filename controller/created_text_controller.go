package controller

import (
	"ai-typing/usecase"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ICreatedTextController interface {
	GetAllCreatedTexts(context echo.Context) error
	FindByGameId(context echo.Context) error
}

type createdTextController struct {
	createdTextUsecase usecase.ICreatedTextUsecase
}

func NewCreatedTextController(createdTextUsecase usecase.ICreatedTextUsecase) ICreatedTextController {
	return &createdTextController{createdTextUsecase}
}

func (createdTextController *createdTextController) GetAllCreatedTexts(context echo.Context) error {
	createdTexts, err := createdTextController.createdTextUsecase.GetAllCreatedTexts()
	if err != nil {
		return context.JSON(500, err.Error())
	}
	return context.JSON(200, createdTexts)
}

func (createdTextController *createdTextController) FindByGameId(context echo.Context) error {
	gameId := context.Param("gameId")
	createdTextsRes, err := createdTextController.createdTextUsecase.FindByGameId(gameId)
	if err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, createdTextsRes)
}
