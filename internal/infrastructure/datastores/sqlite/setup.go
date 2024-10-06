// Package sqlite provides the database operations for the todo endpoint.
package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// New creates a new database connection
func New(filename string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// NewMemory creates a new in-memory database connection
func NewMemory() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
