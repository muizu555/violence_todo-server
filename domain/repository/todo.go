package repository

import "violence_todo/domain/model"

type ITodoRepository interface {
	GetAllTodos(todos *[]model.Todo, userId uint) error
	GetTodoById(todo *model.Todo, userId uint, todoId uint) error
	CreateTodo(todo *model.Todo) error
	UpdateTodo(todo *model.Todo, userId uint, todoId uint) error
	DeleteTodo(userId uint, todoId uint) error
}
