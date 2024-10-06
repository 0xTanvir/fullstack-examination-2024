package sqlite

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/zuu-development/fullstack-examination-2024/internal/domain/bo"
	"github.com/zuu-development/fullstack-examination-2024/internal/domain/definition"
	"gorm.io/gorm"
)

type todo struct {
	db *gorm.DB
}

// NewTodo returns a new instance of the todo repository.
func NewTodo(db *gorm.DB) definition.TodoRepository {
	return &todo{
		db: db,
	}
}

func (td *todo) Create(t *bo.Todo) error {
	if err := td.db.Create(t).Error; err != nil {
		return err
	}
	return nil
}

func (td *todo) Update(t *bo.Todo) error {
	if err := td.db.Save(t).Error; err != nil {
		return err
	}
	return nil
}

func (td *todo) Delete(id int) error {
	result := td.db.Where("id = ?", id).Delete(&bo.Todo{})
	if result.RowsAffected == 0 {
		return bo.ErrNotFound
	}
	if result.Error != nil {
		return result.Error
	}
	log.Info("Deleted todo with id: ", id)
	return nil
}

func (td *todo) Find(id int) (*bo.Todo, error) {
	var todo *bo.Todo
	err := td.db.Where("id = ?", id).Take(&todo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, bo.ErrNotFound
		}
		return nil, err
	}
	return todo, nil
}

func (td *todo) FindAll(tq *bo.TodoQuery) ([]*bo.Todo, error) {
	var todos []*bo.Todo
	query := td.db

	// Apply filters if provided
	if tq.Task != "" {
		query = query.Where("task LIKE ?", "%"+tq.Task+"%")
	}
	if tq.Status != "" {
		query = query.Where("status = ?", tq.Status)
	}

	// Apply ordering if provided else default to DESC
	if tq.Order != "" {
		query = query.Order(fmt.Sprintf("priority %s", tq.Order))
	} else {
		query = query.Order("priority DESC")
	}

	// var todos []*bo.Todo
	err := query.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}
