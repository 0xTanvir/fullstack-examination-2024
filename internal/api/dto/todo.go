// Package dto contains the data transfer object for the application.
package dto

import "github.com/zuu-development/fullstack-examination-2024/internal/domain/bo"

const (
	// PriorityLow is the low priority task.
	PriorityLow = Priority("low")
	// PriorityMedium is the medium priority task.
	PriorityMedium = Priority("medium")
	// PriorityHigh is the high priority task.
	PriorityHigh = Priority("high")

	// Created is the status for a created task.
	Created = Status("created")
	// Processing is the status for a processing task.
	Processing = Status("processing")
	// Done is the status for a done task.
	Done = Status("done")
)

// Status is the status of the task.
type Status string

// Priority is the priority of the task.
type Priority string

// TodoQuery is the query parameter for finding all todos
type TodoQuery struct {
	// Limit  int    `query:"limit,default=20" json:"limit,omitempty" binding:"min=1"`
	// Offset int    `query:"offset" json:"offset,omitempty" binding:"omitempty,min=0"`
	Task   string `query:"task" json:"task,omitempty"`
	Status string `query:"status" json:"status,omitempty"`
	// Sort   string `query:"sort" json:"sort,omitempty"`
	Order string `query:"order" json:"order,omitempty"`
}

// Model creates a new todo query model from the request.
func (tq TodoQuery) Model() *bo.TodoQuery {
	return &bo.TodoQuery{
		Task:   tq.Task,
		Status: tq.Status,
		Order:  tq.Order,
	}
}

// Todo is the model for the todo endpoint.
type Todo struct {
	ID       int    `json:"id"`
	Task     string `json:"task"`
	Status   string `json:"status"`
	Priority string `json:"priority"`
}

// NewTodoFromBo creates a new Todo from a business object.
func NewTodoFromBo(t *bo.Todo) *Todo {
	return &Todo{
		ID:       t.ID,
		Task:     t.Task,
		Status:   t.Status,
		Priority: string(PriorityMapReverse[t.Priority]),
	}
}

// CreateTodoRequest is the request parameter for creating a new todo
type CreateTodoRequest struct {
	Task     string   `json:"task" validate:"required"`
	Priority Priority `json:"priority" validate:"required"`
}

// Model creates a new todo model from the request.
func (tr CreateTodoRequest) Model() *bo.Todo {
	return &bo.Todo{
		Task:     tr.Task,
		Priority: PriorityMap[tr.Priority],
		Status:   string(Created),
	}
}

// UpdateTodoRequest is the request parameter for updating a todo
type UpdateTodoRequest struct {
	UpdateTodoRequestBody
	IDWrapper
}

// UpdateTodoRequestBody is the request body for updating a todo
type UpdateTodoRequestBody struct {
	Task     string   `json:"task,omitempty"`
	Status   Status   `json:"status,omitempty"`
	Priority Priority `json:"priority,omitempty"`
}

// Model creates a new todo model from the request.
func (ur UpdateTodoRequest) Model() *bo.Todo {
	return &bo.Todo{
		ID:       ur.ID,
		Task:     ur.Task,
		Status:   string(ur.Status),
		Priority: PriorityMap[ur.Priority],
	}
}

// ValidPriorities is a map of task priority.
var ValidPriorities = map[Priority]bool{
	PriorityLow:    true,
	PriorityMedium: true,
	PriorityHigh:   true,
}

// PriorityMap is a map of task priority.
var PriorityMap = map[Priority]int{
	PriorityLow:    1,
	PriorityMedium: 2,
	PriorityHigh:   3,
}

// PriorityMapReverse is a map of task priority.
var PriorityMapReverse = map[int]Priority{
	1: PriorityLow,
	2: PriorityMedium,
	3: PriorityHigh,
}

// StatusMap is a map of task status.
var StatusMap = map[Status]bool{
	Created:    true,
	Processing: true,
	Done:       true,
}
