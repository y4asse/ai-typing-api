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
	GetCreatedText(context echo.Context) error
	GetLatestGames(context echo.Context) error
	GetTotalGameCount(context echo.Context) error
	UpdateGameScore(context echo.Context) error
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

func (gameController *gameController) GetCreatedText(context echo.Context) error {
	type RequestBody struct {
		GameId string `json:"game_id"`
	}
	var requestBody RequestBody
	if err := context.Bind(&requestBody); err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	gameId := requestBody.GameId
	createdTextsRes, err := gameController.gameUseCase.GetCreatedText(gameId)
	if err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, createdTextsRes)
}

func (gameController *gameController) GetLatestGames(context echo.Context) error {
	type RequestBody struct {
		Offset int `json:"offset"`
	}
	var requestBody RequestBody
	if err := context.Bind(&requestBody); err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	offset := requestBody.Offset

	gamesRes, err := gameController.gameUseCase.GetLatestGames(offset)
	if err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, gamesRes)
}

func (gameController *gameController) GetTotalGameCount(context echo.Context) error {
	totalGameCount, err := gameController.gameUseCase.GetTotalGameCount()
	if err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, totalGameCount)
}

func (gameController *gameController) UpdateGameScore(context echo.Context) error {
	gameId := context.Param("id")

	game := model.Game{}
	if err := context.Bind(&game); err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	score := game.Score

	err := gameController.gameUseCase.UpdateGameScore(score, gameId)
	if err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, "success")
}
