package router

import (
	"ai-typing/controller"
	"ai-typing/middleWare"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	openaiController controller.IOpenaiController,
	gameController controller.IGameController,
	createdTextController controller.ICreatedTextController,
	likeController controller.IlikeController,
	userController controller.IUserController,
	batchController controller.IBatchController,
) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("FRONT_URL"), os.Getenv("FRONT_DEV_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken, echo.HeaderAuthorization},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	// e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	// 	CookiePath:     "/",
	// 	CookieDomain:   os.Getenv("API_DOMAIN"),
	// 	CookieHTTPOnly: true,
	// 	CookieSameSite: http.SameSiteDefaultMode,
	// }))

	//csrf
	// e.GET("/csrf", userController.CsrfToken)

	//openai
	e.GET("/aiText", openaiController.GetAiText)
	e.POST("/analyse", openaiController.Analyse)

	//game
	// e.GET("/game", gameController.GetAllGame)
	e.PUT("/gameScore/:id", gameController.UpdateGameScore)
	e.GET("/gameRanking", gameController.GetGameRanking)
	e.POST("/latestGames", gameController.GetLatestGames)
	e.GET("/totalGameCount", gameController.GetTotalGameCount)
	e.GET("/games", gameController.GetAllByUserId, middleWare.Auth())

	//createdText
	e.GET("/createdText/:gameId", createdTextController.FindByGameId)
	e.GET("/createdText", createdTextController.GetAllCreatedTexts)

	//like
	e.GET("/likes", likeController.FetchAll)
	e.GET("/likes/:gameId", likeController.FetchAllByGameId)
	e.GET("/likeNum/:gameId", likeController.GetNumByGameId)
	e.POST("/like", likeController.Create)
	e.DELETE("/like/:gameId", likeController.Delete)
	e.GET("likeCountRanking", likeController.GetCountGroupByGameIdOrder)
	e.GET("likedGameIdCount", likeController.GetGameIdCount)

	//user
	e.GET("/gameHistory", gameController.GetGameHistory, middleWare.Auth())
	e.GET("/gameDetail/:gameId", gameController.GetDetail, middleWare.Auth())
	e.POST("/game", gameController.CreateGame, middleWare.AuthAllowGuest())
	e.GET("user", userController.GetUser, middleWare.Auth())
	e.PUT("user", userController.Update, middleWare.Auth())

	//batch
	e.GET("/batches", batchController.GetAllByUserId, middleWare.Auth())

	return e
}
