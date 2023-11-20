package repositories

import (
	"github.com/KKogaa/todo-scheduler/core"
)

type TaskRepository interface {
	FindAll() ([]core.Task, error)
	FindToday() ([]core.Task, error)
	Create(core.Task) error
	Modify(core.Task) error
	Delete(core.Task) error
}
