package handler

import (
	"violence_todo/usecase"

	"github.com/labstack/echo/v4"
)

type ITodoHandler interface {
	Create(e echo.Context)
	Update(c echo.Context)
	Delete(c echo.Context)
	Find(c echo.Context)
	FindAll(c echo.Context)
}

type todoHandler struct {
	todoUsecase usecase.ITodoUsecase
}

func NewTodo(tu usecase.ITodoUsecase) ITodoHandler {

}
