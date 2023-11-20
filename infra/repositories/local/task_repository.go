package local

import (
	"github.com/KKogaa/todo-scheduler/core"
)

type TaskRepository struct {
}

func NewTaskRepository() TaskRepository {
	return TaskRepository{}
}

func (t TaskRepository) FindAll() ([]core.Task, error) {
	return nil, nil
}

func (t TaskRepository) FindToday() ([]core.Task, error) {
	return nil, nil
}

func (t TaskRepository) Create(task core.Task) error {
	return nil
}

func (t TaskRepository) Modify(task core.Task) error {
	return nil
}

func (t TaskRepository) Delete(task core.Task) error {
	return nil
}
