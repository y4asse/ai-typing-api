package controller

import (
	"ai-typing/model"
	"ai-typing/usecase"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IOpenaiController interface {
	GetAiText(c echo.Context) error
	Analyse(c echo.Context) error
}

type openaiController struct {
	ou usecase.IOpenaiUsecase
}

func NewOpenaiController(ou usecase.IOpenaiUsecase) IOpenaiController {
	return &openaiController{ou}
}

func (oc *openaiController) GetAiText(c echo.Context) error {
	//query paramからデータを取り出す
	thema := c.QueryParam("thema")
	openaiRes, err := oc.ou.GetAiText(thema)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, openaiRes)
}

func (oc *openaiController) Analyse(c echo.Context) error {
	var requestBody model.AnalyseRequest
	if err := c.Bind(&requestBody); err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	analyseRes, err := oc.ou.Analyse(requestBody)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, analyseRes)
}
