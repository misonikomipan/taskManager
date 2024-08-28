package main

import (
	"os"
	"testing"
	"time"
)

func TestFileExists(t *testing.T) {
	filename := "test_tasks.json"
	if fileExists(filename) {
		t.Errorf("Expected fileExists to return false for non-existent file, got true")
	}
	// ファイルを作成
	if _, err := os.Create(filename); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filename)

	if !fileExists(filename) {
		t.Errorf("Expected fileExists to return true for existing file, got false")
	}
}

func TestLoadTasks(t *testing.T) {
	// ダミータスクデータを作成
	tasks := []*Task{
		{
			ID:        1,
			Name:      "Test Task 1",
			CreatedAt: time.Now(),
			Completed: false,
		},
		{
			ID:        2,
			Name:      "Test Task 2",
			CreatedAt: time.Now().Add(-time.Hour * 24), // 1日前
			Completed: true,
		},
	}

	// テスト用のファイル名
	testFileName := "test_tasks.json"

	// テストファイルにタスクを保存
	err := SaveTasks(testFileName, tasks)
	if err != nil {
		t.Fatalf("SaveTasks: Expected no error, got %v", err)
	}

	// LoadTasks関数をテスト
	loadedTasks, err := LoadTasks(testFileName)
	if err != nil {
		t.Fatalf("LoadTasks: Expected no error, got %v", err)
	}

	// 保存したタスクとロードしたタスクが同じかどうかを確認
	if len(loadedTasks) != len(tasks) {
		t.Fatalf("LoadTasks: Expected %d tasks, got %d", len(tasks), len(loadedTasks))
	}

	for i, task := range tasks {
		if task.ID != loadedTasks[i].ID || task.Name != loadedTasks[i].Name || task.Completed != loadedTasks[i].Completed {
			t.Errorf("LoadTasks: Task %d does not match. Expected %+v, got %+v", i, task, loadedTasks[i])
		}
	}

	// テストが終了したらテストファイルを削除
	os.Remove(testFileName)
}

func TestSaveTasks(t *testing.T) {
	tasks := []*Task{
		{
			ID:        1,
			Name:      "Test Task 1",
			CreatedAt: time.Now(),
			Completed: false,
		},
		{
			ID:        2,
			Name:      "Test Task 2",
			CreatedAt: time.Now().Add(-time.Hour * 24), // 1日前
			Completed: true,
		},
	}
	testFileName := "test_tasks.json"
	// SaveTasks関数をテスト
	err := SaveTasks(testFileName, tasks)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// ファイルが正しく保存されたかを確認するために、LoadTasksを利用
	loadedTasks, err := LoadTasks(testFileName)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// 保存したタスクとロードしたタスクが同じかどうかを確認
	if len(loadedTasks) != len(tasks) {
		t.Fatalf("Expected %d tasks, got %d", len(tasks), len(loadedTasks))
	}

	for i, task := range tasks {
		if task.ID != loadedTasks[i].ID || task.Name != loadedTasks[i].Name || task.Completed != loadedTasks[i].Completed {
			t.Errorf("Task %d does not match. Expected %+v, got %+v", i, task, loadedTasks[i])
		}
	}

	// テストが終了したらテストファイルを削除
	os.Remove(testFileName)
}
