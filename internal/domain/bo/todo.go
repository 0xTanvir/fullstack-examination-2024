// Package bo Business object for the applications.
package bo

import "time"

// Todo is the model for the todo endpoint.
type Todo struct {
	ID        int `gorm:"primaryKey"`
	Task      string
	Status    string
	Priority  int       // 1: low, 2: medium, 3: high
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// TodoQuery is the query parameter for the todo endpoint.
type TodoQuery struct {
	// Limit  int
	// Offset int
	Task   string
	Status string
	// Sort   string
	Order string
}
