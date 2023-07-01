package router

import (
	"ai-typing/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(openaiController controller.IOpenaiController, gameController controller.IGameController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FRONT_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.POST("/aiText", openaiController.GetAiText)
	e.POST("/game", gameController.CreateGame)
	e.POST("/gameHistory", gameController.GetGameHistory)
	e.POST("/createdText", gameController.GetCreatedText)
	e.GET("/game", gameController.GetAllGame)
	e.GET("/gameRanking", gameController.GetGameRanking)
	return e
}
