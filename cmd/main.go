package main

import (
	"violence_todo/infra"
	"violence_todo/usecase"
)

func main() {
	db := infra.NewDB()
	todoRepository := infra.NewTodo(db)
	todoUsecase := usecase.NewTodo(todoRepository)
	todoHandler := hander

}
