// Package definition provides all the definition for the application.
package definition

import "github.com/zuu-development/fullstack-examination-2024/internal/api/dto"

// TodoService is the service for the todo endpoint.
type TodoService interface {
	Create(ctr dto.CreateTodoRequest) (*dto.Todo, error)
	Update(utr dto.UpdateTodoRequest) (*dto.Todo, error)
	Delete(id int) error
	Find(id int) (*dto.Todo, error)
	FindAll(tq dto.TodoQuery) ([]*dto.Todo, error)
}
