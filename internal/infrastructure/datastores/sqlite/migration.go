package sqlite

import (
	"github.com/zuu-development/fullstack-examination-2024/internal/domain/bo"
	"gorm.io/gorm"
)

// Migrate runs the auto-migration for the database
func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&bo.Todo{}); err != nil {
		return err
	}
	return nil
}
