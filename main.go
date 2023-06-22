// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"ai-typing/controller"
	"ai-typing/db"
	"ai-typing/repository"

	"ai-typing/router"
	"ai-typing/usecase"
	"os"
)

func main() {
	db := db.NewDB()
	gameRepository := repository.NewGameRepository(db)
	createdTextRepository := repository.NewCreatedTextRepository(db)
	gameUsecase := usecase.NewGameUsecase(gameRepository)
	createdTextUsecase := usecase.NewCreatedTextUsecase(createdTextRepository)
	gameController := controller.NewGameController(gameUsecase, createdTextUsecase)
	openaiUsecase := usecase.NewOpenaiUsecase()
	openaiController := controller.NewOpenaiController(openaiUsecase)
	e := router.NewRouter(openaiController, gameController)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
