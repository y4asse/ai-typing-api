package migrate

import (
	"fmt"

	"ai-typing/db"
	"ai-typing/model"
)

func Migrate() {
	fmt.Println("Migrating...")
	dbConnection := db.NewDB()
	defer fmt.Println("Successfully Migrated!!!!!")
	defer db.CloseDB(dbConnection)
	dbConnection.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	dbConnection.AutoMigrate(&model.User{}, &model.Comment{}, &model.CreatedText{}, &model.Mode{}, &model.Game{}, &model.Like{}, &model.PostedText{})
}
