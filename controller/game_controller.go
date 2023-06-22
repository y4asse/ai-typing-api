package controller

import (
	"ai-typing/model"
	"ai-typing/usecase"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type IGameController interface {
	CreateGame(context echo.Context) error
}

type gameController struct {
	gameUseCase        usecase.IGameUsecase
	createdTextUsecase usecase.ICreatedTextUsecase
}

func NewGameController(gameUseCase usecase.IGameUsecase, createdTextUsecase usecase.ICreatedTextUsecase) IGameController {
	return &gameController{gameUseCase, createdTextUsecase}
}

func (gameController *gameController) CreateGame(context echo.Context) error {
	tempGame := model.TempGame{}
	if err := context.Bind(&tempGame); err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	game := model.Game{}
	game.ID = uuid.NewString()
	game.InputedThema = tempGame.InputedThema
	game.ModeId = tempGame.ModeId
	game.Score = tempGame.Score

	createdText := model.CreatedText{}
	createdText.ID = uuid.NewString()
	createdText.Text = tempGame.Text
	createdText.Hiragana = tempGame.Hiragana
	createdText.GameId = game.ID

	gameRes, err := gameController.gameUseCase.CreateGame(game)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	createdTextRes, err := gameController.createdTextUsecase.CreateCreatedText(createdText)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(createdTextRes)
	return context.JSON(http.StatusCreated, gameRes)

}
