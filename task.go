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

var taskIDCounter = 0

func NewTask(name string) *Task {
	taskIDCounter++
	return &Task{
		ID:        taskIDCounter,
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
