package router

import (
	"ai-typing/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(oc controller.IOpenaiController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	e.POST("/aiText", oc.GetAiText)
	return e
}
