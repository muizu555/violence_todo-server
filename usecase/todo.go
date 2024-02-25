package usecase

import (
	"violence_todo/domain/model"
	"violence_todo/domain/repository"
)

type ITodoUsecase interface {
	GetAllTodos(userId uint) ([]model.TodoResponse, error)
	GetTodoById(userId uint, todoId uint) (model.TodoResponse, error)
	CreateTodo(todo model.Todo) (model.TodoResponse, error)
	UpdateTodo(todo model.Todo, userId uint, todoId uint) (model.TodoResponse, error)
	DeleteTodo(userId uint, todoId uint) error
}

type todoUsecase struct {
	tr repository.ITodoRepository
}

func NewTodo(tr repository.ITodoRepository) ITodoUsecase { //これは依存性注入(DI)に使うはず。。。。
	return &todoUsecase{tr}
}

func (tu *todoUsecase) GetAllTodos(userId uint) ([]model.TodoResponse, error) {
	todos := []model.Todo{} //一旦全部初期化で
	if err := tu.tr.GetAllTodos(&todos, userId); err != nil {
		return nil, err //ここはなぜ、nilが使えるんだっけ？
	}
	resTodos := []model.TodoResponse{}
	for _, v := range todos {
		t := model.TodoResponse{
			ID:        v.ID,
			Content:   v.Content,
			Checked:   v.Checked,
			CreatedAt: v.CreatedAt,
		}
		resTodos = append(resTodos, t)
	}
	return resTodos, nil
}

func (tu *todoUsecase) GetTodoById(userId uint, todoId uint) (model.TodoResponse, error) {
	todo := model.Todo{}
	if err := tu.tr.GetTodoById(&todo, userId, todoId); err != nil {
		return model.TodoResponse{}, err
	}
	resTodo := model.TodoResponse{
		ID:        todo.ID,
		Content:   todo.Content,
		Checked:   todo.Checked,
		CreatedAt: todo.CreatedAt,
	}
	return resTodo, nil
}

func (tu *todoUsecase) CreateTodo(todo model.Todo) (model.TodoResponse, error) {
	if err := tu.tr.CreateTodo(&todo); err != nil {
		return model.TodoResponse{}, err
	}
	resTodo := model.TodoResponse{
		ID:        todo.ID,
		Content:   todo.Content,
		Checked:   todo.Checked,
		CreatedAt: todo.CreatedAt,
	}
	return resTodo, nil
}

func (tu *todoUsecase) UpdateTodo(todo model.Todo, userId uint, todoId uint) (model.TodoResponse, error) {
	if err := tu.tr.UpdateTodo(&todo, userId, todoId); err != nil {
		return model.TodoResponse{}, err
	}
	resTodo := model.TodoResponse{
		ID:        todo.ID,
		Content:   todo.Content,
		Checked:   todo.Checked,
		CreatedAt: todo.CreatedAt,
	}
	return resTodo, nil
}

func (tu *todoUsecase) DeleteTodo(userId uint, todoId uint) error {
	if err := tu.tr.DeleteTodo(userId, todoId); err != nil {
		return err
	}
	return nil
}
