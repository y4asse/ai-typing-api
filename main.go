// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"ai-typing/controller"
	"ai-typing/db"
	"ai-typing/router"
	"ai-typing/usecase"
)

func main() {
	db.NewDB()
	openaiUsecase := usecase.NewOpenaiUsecase()
	openaiController := controller.NewOpenaiController(openaiUsecase)
	e := router.NewRouter(openaiController)
	e.Logger.Fatal(e.Start(":8080"))
}
