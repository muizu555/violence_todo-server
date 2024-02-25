package main

import (
	"violence_todo/handler"
	"violence_todo/infra"
	"violence_todo/usecase"
)

func main() {
	db := infra.NewDB()
	todoRepository := infra.NewTodo(db)
	todoUsecase := usecase.NewTodo(todoRepository)
	todoHandler := handler.NewTodo(todoUsecase)

}
