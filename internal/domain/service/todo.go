// Package service provides the business logic for the todo endpoint.
package service

import (
	"github.com/zuu-development/fullstack-examination-2024/internal/api/dto"
	"github.com/zuu-development/fullstack-examination-2024/internal/domain/definition"
)

type todo struct {
	todoRepository definition.TodoRepository
}

// NewTodo creates a new Todo service.
func NewTodo(r definition.TodoRepository) definition.TodoService {
	return &todo{r}
}

func (t *todo) Create(ctr dto.CreateTodoRequest) (*dto.Todo, error) {
	todo := ctr.Model()
	if err := t.todoRepository.Create(todo); err != nil {
		return nil, err
	}
	return dto.NewTodoFromBo(todo), nil
}

func (t *todo) Update(utr dto.UpdateTodoRequest) (*dto.Todo, error) {
	todo := utr.Model()
	// 現在の値を取得
	currentTodo, err := t.todoRepository.Find(todo.ID)
	if err != nil {
		return nil, err
	}
	// 空文字列の場合、現在の値を使用
	if todo.Task == "" {
		todo.Task = currentTodo.Task
	}
	if todo.Priority == 0 {
		todo.Priority = currentTodo.Priority
	}
	if todo.Status == "" {
		todo.Status = currentTodo.Status
	}

	if err := t.todoRepository.Update(todo); err != nil {
		return nil, err
	}
	return dto.NewTodoFromBo(todo), nil
}

func (t *todo) Delete(id int) error {
	if err := t.todoRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (t *todo) Find(id int) (*dto.Todo, error) {
	todo, err := t.todoRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return dto.NewTodoFromBo(todo), nil
}

func (t *todo) FindAll(tq dto.TodoQuery) ([]*dto.Todo, error) {
	query := tq.Model()

	boTodos, err := t.todoRepository.FindAll(query)
	if err != nil {
		return nil, err
	}

	todos := []*dto.Todo{}
	for _, todo := range boTodos {
		todos = append(todos, dto.NewTodoFromBo(todo))
	}

	return todos, nil
}
