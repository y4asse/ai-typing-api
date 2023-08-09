// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"ai-typing/controller"
	"ai-typing/db"
	"ai-typing/migrate"
	"ai-typing/repository"

	"ai-typing/router"
	"ai-typing/usecase"
	"os"
)

func main() {
	migrate.Migrate()
	db := db.NewDB()
	//repository
	gameRepository := repository.NewGameRepository(db)
	createdTextRepository := repository.NewCreatedTextRepository(db)
	likeRepository := repository.NewLikeRepository(db)

	//usecase
	gameUsecase := usecase.NewGameUsecase(gameRepository, createdTextRepository)
	createdTextUsecase := usecase.NewCreatedTextUsecase(createdTextRepository)
	likeUsecase := usecase.NewLikeUsecase(likeRepository, gameRepository)

	//controller
	likeController := controller.NewLikeController(likeUsecase)
	createTextController := controller.NewCreatedTextController(createdTextUsecase)
	gameController := controller.NewGameController(gameUsecase, createdTextUsecase)
	openaiUsecase := usecase.NewOpenaiUsecase()
	openaiController := controller.NewOpenaiController(openaiUsecase)
	e := router.NewRouter(openaiController, gameController, createTextController, likeController)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
