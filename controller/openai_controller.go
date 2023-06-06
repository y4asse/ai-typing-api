package controller

import (
	"ai-typing/model"
	"ai-typing/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IOpenaiController interface {
	GetAiText(c echo.Context) error
}

type openaiController struct {
	ou usecase.IOpenaiUsecase
}

func NewOpenaiController(ou usecase.IOpenaiUsecase) IOpenaiController {
	return &openaiController{ou}
}

func (oc *openaiController) GetAiText(c echo.Context) error {
	var requestBody model.AiTextRequest
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	thema := requestBody.Thema
	openaiRes, err := oc.ou.GetAiText(thema)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, openaiRes)
}
