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
	GetGameRanking(context echo.Context) error
	GetGameHistory(context echo.Context) error
	GetAllGame(context echo.Context) error
}

type gameController struct {
	gameUseCase        usecase.IGameUsecase
	createdTextUsecase usecase.ICreatedTextUsecase
}

func NewGameController(gameUseCase usecase.IGameUsecase, createdTextUsecase usecase.ICreatedTextUsecase) IGameController {
	return &gameController{gameUseCase, createdTextUsecase}
}

func (gameController *gameController) CreateGame(context echo.Context) error {
	gameBody := model.GameBody{}
	if err := context.Bind(&gameBody); err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	game := model.Game{
		ID:           uuid.NewString(),
		InputedThema: gameBody.InputedThema,
		ModeId:       gameBody.ModeId,
		Score:        gameBody.Score,
		UserId:       gameBody.UserId,
	}
	for i := range gameBody.Text {
		createdText := model.CreatedText{
			ID:       uuid.NewString(),
			Text:     gameBody.Text[i],
			Hiragana: gameBody.Hiragana[i],
			GameId:   game.ID,
		}
		_, err := gameController.createdTextUsecase.CreateCreatedText(createdText)
		if err != nil {
			fmt.Println(err.Error())
			return context.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	gameRes, err := gameController.gameUseCase.CreateGame(game)
	if err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusCreated, gameRes)
}

func (gameController *gameController) GetGameRanking(context echo.Context) error {
	gamesRes, err := gameController.gameUseCase.GetGameRanking()
	if err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, gamesRes)
}

func (gameController *gameController) GetGameHistory(context echo.Context) error {
	type RequestBody struct {
		UserId string `json:"user_id"`
	}
	var requestBody RequestBody
	if err := context.Bind(&requestBody); err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	userId := requestBody.UserId
	gamesRes, err := gameController.gameUseCase.GetGameHistory(userId)
	if err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, gamesRes)
}

func (gameController *gameController) GetAllGame(context echo.Context) error {
	gamesRes, err := gameController.gameUseCase.GetAllGame()
	if err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, gamesRes)
}
