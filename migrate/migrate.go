package main

import (
	"fmt"
	"violence_todo/db"
	"violence_todo/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("success migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Todo{}, &model.Comment{})
}
