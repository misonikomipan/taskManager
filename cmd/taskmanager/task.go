package main

import (
	"errors"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Completed bool      `json:"completed"`
}

func NewTask(id int, name string) *Task {
	return &Task{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
		Completed: false,
	}
}

func (t *Task) ToggleCompleted() {
	t.Completed = !t.Completed
}

func (t *Task) UpdateName(name string) error {
	// 空文字列を渡した場合
	if name == "" {
		return errors.New("name can't be empty")
	}
	t.Name = name
	return nil
}

func getNextTaskID(tasks []*Task) int {
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}
