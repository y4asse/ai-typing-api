package router

import (
	"ai-typing/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(oc controller.IOpenaiController) *echo.Echo {
	e := echo.New()
	e.POST("/aiText", oc.GetAiText)
	return e
}
