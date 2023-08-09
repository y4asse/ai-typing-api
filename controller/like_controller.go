package controller

import (
	"ai-typing/model"
	"ai-typing/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IlikeController interface {
	FetchAll(context echo.Context) error
	Create(context echo.Context) error
	Delete(context echo.Context) error
	FetchAllByGameId(context echo.Context) error
	GetNumByGameId(context echo.Context) error
	GetCountGroupByGameIdOrder(context echo.Context) error
}

type likeController struct {
	likeUsecase usecase.ILikeUsecase
}

func NewLikeController(likeUsecase usecase.ILikeUsecase) IlikeController {
	return &likeController{likeUsecase}
}

func (likeController *likeController) FetchAll(context echo.Context) error {
	likes, err := likeController.likeUsecase.FetchAll()
	if err != nil {
		return context.JSON(500, err.Error())
	}
	return context.JSON(200, likes)
}

func (likeController *likeController) Create(context echo.Context) error {
	like := model.Like{}
	if err := context.Bind(&like); err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	err := likeController.likeUsecase.Create(like)
	if err != nil {
		return context.JSON(500, err.Error())
	}
	return context.JSON(200, like)
}

func (likeController *likeController) Delete(context echo.Context) error {
	gameId := context.Param("gameId")
	err := likeController.likeUsecase.Delete(gameId)
	if err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, gameId)
}

func (likeController *likeController) FetchAllByGameId(context echo.Context) error {
	gameId := context.Param("gameId")
	likes, err := likeController.likeUsecase.FetchAllByGameId(gameId)
	if err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, likes)
}

func (likeController *likeController) GetNumByGameId(context echo.Context) error {
	gameId := context.Param("gameId")
	num, err := likeController.likeUsecase.GetNumByGameId(gameId)
	if err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, num)
}

func (likeController *likeController) GetCountGroupByGameIdOrder(context echo.Context) error {
	offset, err := strconv.Atoi(context.QueryParam("offset"))
	if err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	limit, err := strconv.Atoi(context.QueryParam("limit"))
	if err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	gamesWithCount, err := likeController.likeUsecase.GetCountGroupByGameIdOrder(offset, limit)
	if err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, gamesWithCount)
}
