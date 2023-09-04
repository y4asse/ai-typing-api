// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"ai-typing/controller"
	"ai-typing/db"
	"ai-typing/migrate"
	"ai-typing/repository"
	"ai-typing/validator"

	"ai-typing/router"
	"ai-typing/usecase"
	"os"
)

func main() {
	migrate.Migrate()
	db := db.NewDB()

	//validator
	aiTextValidator := validator.NewAitextValidator()

	//repository
	gameRepository := repository.NewGameRepository(db)
	createdTextRepository := repository.NewCreatedTextRepository(db)
	likeRepository := repository.NewLikeRepository(db)
	userRepository := repository.NewUserRepository(db)

	//usecase
	gameUsecase := usecase.NewGameUsecase(gameRepository, createdTextRepository)
	createdTextUsecase := usecase.NewCreatedTextUsecase(createdTextRepository)
	likeUsecase := usecase.NewLikeUsecase(likeRepository, gameRepository)
	userusecase := usecase.NewUserUsecase(userRepository)

	//controller
	likeController := controller.NewLikeController(likeUsecase)
	createTextController := controller.NewCreatedTextController(createdTextUsecase)
	gameController := controller.NewGameController(gameUsecase, createdTextUsecase)
	openaiUsecase := usecase.NewOpenaiUsecase(aiTextValidator)
	openaiController := controller.NewOpenaiController(openaiUsecase)
	userController := controller.NewUserController(userusecase)

	e := router.NewRouter(openaiController, gameController, createTextController, likeController, userController)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
