package main

import (
	"fmt"

	"ai-typing/db"
	"ai-typing/model"
)

func main() {
	fmt.Println("Migrating...")
	dbConnection := db.NewDB()
	defer fmt.Println("Successfully Migrated!!!!!")
	defer db.CloseDB(dbConnection)
	dbConnection.AutoMigrate(&model.User{}, &model.Comment{}, &model.CreatedText{}, &model.Difficulty{}, &model.Game{}, &model.Like{}, &model.PostedText{})
}
