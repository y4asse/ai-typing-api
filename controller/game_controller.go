package controller

import (
	"ai-typing/model"
	"ai-typing/usecase"
	"fmt"
	"net/http"
	"strconv"

	"firebase.google.com/go/auth"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type IGameController interface {
	CreateGame(context echo.Context) error
	GetGameRanking(context echo.Context) error
	GetGameHistory(context echo.Context) error
	GetAllGame(context echo.Context) error
	GetLatestGames(context echo.Context) error
	GetTotalGameCount(context echo.Context) error
	UpdateGameScore(context echo.Context) error
	GetAllByUserId(context echo.Context) error
	GetDetail(context echo.Context) error
}

type gameController struct {
	gameUseCase        usecase.IGameUsecase
	createdTextUsecase usecase.ICreatedTextUsecase
}

func NewGameController(gameUseCase usecase.IGameUsecase, createdTextUsecase usecase.ICreatedTextUsecase) IGameController {
	return &gameController{gameUseCase, createdTextUsecase}
}

func (gameController *gameController) CreateGame(context echo.Context) error {
	token := context.Get("token").(*auth.Token)
	claims := token.Claims
	uid, _ := claims["user_id"].(string)
	gameBody := model.GameBody{}
	if err := context.Bind(&gameBody); err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	game := model.Game{
		ID:             uuid.NewString(),
		InputedThema:   gameBody.InputedThema,
		ModeId:         gameBody.ModeId,
		Score:          gameBody.Score,
		UserId:         uid,
		DisableRanking: gameBody.DisableRanking,
		AiModel:        gameBody.AiModel,
		Detail:         gameBody.Detail,
		TotalKeyCount:  gameBody.TotalKeyCount,
		TotalMissType:  gameBody.TotalMissType,
		TotalTime:      gameBody.TotalTime,
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
	border, _ := strconv.Atoi(context.QueryParam("border"))
	gamesRes, err := gameController.gameUseCase.GetGameRanking(border)
	if err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, gamesRes)
}

func (gameController *gameController) GetGameHistory(context echo.Context) error {
	token := context.Get("token").(*auth.Token)
	claims := token.Claims
	uid, _ := claims["user_id"].(string)
	var limit int
	if context.QueryParam("limit") == "" {
		limit = 10
	} else {
		limit, _ = strconv.Atoi(context.QueryParam("limit"))
	}
	gamesRes, err := gameController.gameUseCase.GetGameHistory(uid, limit)
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
	game := model.Game{}
	if err := context.Bind(&game); err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	game.ID = context.Param("id")
	updateGameResponse, err := gameController.gameUseCase.UpdateGameScore(&game)
	if err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, updateGameResponse)
}

func (gameController *gameController) GetAllByUserId(context echo.Context) error {
	token := context.Get("token").(*auth.Token)
	claims := token.Claims
	uid, _ := claims["user_id"].(string)

	gamesRes, err := gameController.gameUseCase.GetAllByUserId(uid)
	if err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, gamesRes)
}

func (gameController *gameController) GetDetail(context echo.Context) error {
	token := context.Get("token").(*auth.Token)
	claims := token.Claims
	uid, _ := claims["user_id"].(string)
	gameId := context.Param("gameId")
	gameDetail, err := gameController.gameUseCase.GetDetail(gameId, uid)
	if err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, gameDetail)
}
