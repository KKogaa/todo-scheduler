package application

import (
	"time"

	"github.com/KKogaa/todo-scheduler/application/repositories"
	"github.com/KKogaa/todo-scheduler/core"
)

type CreateUsecase struct {
	TaskRepo repositories.TaskRepository
}

func NewCreateUsecase(taskRepo repositories.TaskRepository) *CreateUsecase {
	return &CreateUsecase{
		TaskRepo: taskRepo,
	}
}

func (c CreateUsecase) Execute(description string, name string,
	status core.TaskStatus, estimatedTime time.Duration,
	difficulty core.TaskDifficulty) (core.Task, error) {

	task := core.Task{
		StartTime:     time.Now(),
		Description:   description,
		Name:          name,
		Status:        status,
		EstimatedTime: estimatedTime,
		Difficulty:    difficulty,
	}

	err := c.TaskRepo.Create(task)
	if err != nil {
		return core.Task{}, err
	}

	return task, nil
}
