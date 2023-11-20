package core

import "time"

type TaskStatus int8

const (
	Pending    TaskStatus = 0
	Doing      TaskStatus = 1
	Done       TaskStatus = 2
	Terminated TaskStatus = 3
	Cancelled  TaskStatus = 4
)

type TaskDifficulty int8

const (
	Easy     TaskDifficulty = 0
	Medium   TaskDifficulty = 1
	Hard     TaskDifficulty = 2
	VeryHard TaskDifficulty = 3
	Unknown  TaskDifficulty = 4
)

type Task struct {
	StartTime     time.Time
	Name          string
	Description   string
	Status        TaskStatus
	EstimatedTime time.Duration //eta in idk which metric to use yet, use minutes
	Difficulty    TaskDifficulty
}

func NewTask(name string, description string,
	estimatedTime time.Duration, difficulty TaskDifficulty) *Task {

	return &Task{
		StartTime:     time.Now(),
		Name:          name,
		Description:   description,
		Status:        0,
		EstimatedTime: estimatedTime,
		Difficulty:    difficulty,
	}
}

// check if status time
func (t *Task) changeToDoneByEstimatedTime() error {
	if time.Now().After(t.StartTime.Add(t.EstimatedTime)) {
		//TODO: check setting this using the enum
		t.Status = 2
	}
	return nil
}

func (t Task) getTimeRemaining() time.Duration {
	return time.Now().Sub(t.StartTime)
}
