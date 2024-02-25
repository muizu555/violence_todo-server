package main

import (
	"fmt"

	"violence_todo/domain/model"
	"violence_todo/infra"
)

func main() {
	dbConn := infra.NewDB()
	defer fmt.Println("success migrated")
	defer infra.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Todo{}, &model.Comment{})
}
