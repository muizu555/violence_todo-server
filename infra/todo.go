package infra

import (
	"fmt"
	"violence_todo/domain/model"
	"violence_todo/domain/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Todo struct {
	db *gorm.DB
}

func NewTodo(db *gorm.DB) repository.ITodoRepository {
	return &Todo{db}
}

func (td *Todo) GetAllTodos(todos *[]model.Todo, userId uint) error {
	if err := td.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(todos).Error; err != nil {
		return err
	}
	return nil
}

func (td *Todo) GetTodoById(todo *model.Todo, userId uint, todoId uint) error {
	if err := td.db.Joins("User").Where("user_id=?", userId).First(todo, todoId).Error; err != nil {
		return err
	}
	return nil
}

func (td *Todo) CreateTodo(todo *model.Todo) error {
	if err := td.db.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func (td *Todo) UpdateTodo(todo *model.Todo, userId uint, todoId uint) error {
	result := td.db.Model(todo).Clauses(clause.Returning{}).Where("id=? AND user_id=?", todoId, userId).Update("title", todo.Content)
	if result.Error != nil {
		return result.Error
	}
	//このチェックは、更新が成功したかどうかを確認するために行われています。もしRowsAffectedが1未満であれば、対象のオブジェクトが存在しないことを示し、
	//エラーが返されます。一般的に、更新対象のオブジェクトがデータベースに存在しない場合にこのようなエラーを返すことが一般的です。
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (td *Todo) DeleteTodo(userId uint, todoId uint) error {
	result := td.db.Where("id=? AND user_id=?", todoId, userId).Delete(&model.Todo{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
