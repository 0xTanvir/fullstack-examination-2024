// Package definition provides all the definition for the application.
package definition

import "github.com/zuu-development/fullstack-examination-2024/internal/domain/bo"

// DataStore is the data store definition for the application.
type DataStore struct {
	Todo TodoRepository
}

// TodoRepository is the repository for the todo endpoint.
type TodoRepository interface {
	Create(t *bo.Todo) error
	Delete(id int) error
	Update(t *bo.Todo) error
	Find(id int) (*bo.Todo, error)
	FindAll(*bo.TodoQuery) ([]*bo.Todo, error)
}
