package main

import (
	"testing"
)

func TestToggleCompleted(t *testing.T) {
	task := Task{Name: "Test Task"}
	task.ToggleCompleted()

	if task.Completed != true {
		t.Errorf("Expected Completed to be true, got %v", task.Completed)
	}

	task.ToggleCompleted()

	if task.Completed != false {
		t.Errorf("Expected Completed to be false, got %v", task.Completed)
	}
}

func TestUpdateName(t *testing.T) {
	task := Task{Name: "Old Name"}
	err := task.UpdateName("New Name")
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	if task.Name != "New Name" {
		t.Errorf("Expected Name to be 'New Name', got %v", task.Name)
	}

	// 空文字列を渡した場合
	err = task.UpdateName("")
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}
