package router

import (
	"ai-typing/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	openaiController controller.IOpenaiController,
	gameController controller.IGameController,
	createdTextController controller.ICreatedTextController,
	likeController controller.IlikeController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("FRONT_URL"), os.Getenv("FRONT_DEV_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.POST("/aiText", openaiController.GetAiText)

	//game
	e.GET("/game", gameController.GetAllGame)
	e.POST("/game", gameController.CreateGame)
	e.PUT("/gameScore/:id", gameController.UpdateGameScore)
	e.POST("/gameHistory", gameController.GetGameHistory)
	e.GET("/gameRanking", gameController.GetGameRanking)
	e.POST("/latestGames", gameController.GetLatestGames)
	e.GET("/totalGameCount", gameController.GetTotalGameCount)

	//createdText
	e.GET("/createdText/:gameId", createdTextController.FindByGameId)
	e.GET(("/createdText"), createdTextController.GetAllCreatedTexts)

	//like
	e.GET("/likes", likeController.FetchAll)
	e.GET("/likes/:gameId", likeController.FetchAllByGameId)
	e.GET("/likeNum/:gameId", likeController.GetNumByGameId)
	e.POST("/like", likeController.Create)
	e.DELETE("/like/:gameId", likeController.Delete)
	return e
}
