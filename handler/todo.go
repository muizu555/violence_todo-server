package handler

import (
	"violence_todo/usecase"

	"github.com/labstack/echo/v4"
)

type ITodoHandler interface {
	Create(e echo.Context)
	Update(e echo.Context)
	Delete(e echo.Context)
	Find(e echo.Context)
	FindAll(e echo.Context)
}

type todoHandler struct {
	todoUsecase usecase.ITodoUsecase
}

func NewTodo(tu usecase.ITodoUsecase) ITodoHandler {

}
