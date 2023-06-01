package main

import (
	"fmt"

	"example.com/hello/db"
	"example.com/hello/model"
)

func main() {
	fmt.Println("Migrating...")
	dbConnection := db.NewDB()
	defer fmt.Println("Successfully Migrated!!!!!")
	defer db.CloseDB(dbConnection)
	dbConnection.AutoMigrate(&model.User{})
}
