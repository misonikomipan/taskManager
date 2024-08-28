package main

import "testing"

func TestAddHandler(t *testing.T) {
	// テスト用のタスク名
	taskName := "test task"
	// テスト用のタスクを追加
	id, err := handleAdd(taskName)
	if err != nil {
		t.Fatalf("failed to add task: %v", err)
	}

	// 追加したタスクがtasks.jsonに保存されているか確認
	tasks, err := LoadTasks(tasksFileName)
	if err != nil {
		t.Fatalf("failed to load tasks: %v", err)
	}

	// 追加したタスクがtasksに含まれているか確認
	found := false
	for _, task := range tasks {
		if task.ID == id {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("task not found in tasks")
	}
}

func TestListHandler(t *testing.T) {
	// tasks.jsonにタスクが1つもない状態でlistコマンドを実行
	err := handleList()
	if err != nil {
		t.Fatalf("failed to list tasks: %v", err)
	}

	// テスト用のタスク名
	taskName := "test task"
	// テスト用のタスクを追加
	_, err = handleAdd(taskName)
	if err != nil {
		t.Fatalf("failed to add task: %v", err)
	}

	// 追加したタスクがtasks.jsonに保存されているか確認
	tasks, err := LoadTasks(tasksFileName)
	if err != nil {
		t.Fatalf("failed to load tasks: %v", err)
	}

	// 追加したタスクがtasksに含まれているか確認
	found := false
	for _, task := range tasks {
		if task.Name == taskName {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("task not found in tasks")
	}

	// 追加したタスクが表示されるか確認
	err = handleList()
	if err != nil {
		t.Fatalf("failed to list tasks: %v", err)
	}
}

func TestHandleComplete(t *testing.T) {
	// テスト用のタスク名
	taskName := "test task"
	// テスト用のタスクを追加
	id, err := handleAdd(taskName)
	if err != nil {
		t.Fatalf("failed to add task: %v", err)
	}

	// 追加したタスクがtasks.jsonに保存されているか確認
	tasks, err := LoadTasks(tasksFileName)
	if err != nil {
		t.Fatalf("failed to load tasks: %v", err)
	}

	// 追加したタスクがtasksに含まれているか確認
	found := false
	for _, task := range tasks {
		if task.ID == id {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("task not found in tasks")
	}

	// タスクを完了にする
	err = handleComplete(id)
	if err != nil {
		t.Fatalf("failed to complete task: %v", err)
	}

	// 完了したタスクがtasks.jsonに保存されているか確認
	tasks, err = LoadTasks(tasksFileName)
	if err != nil {
		t.Fatalf("failed to load tasks: %v", err)
	}

	// 完了したタスクがtasksに含まれているか確認
	found = false
	for _, task := range tasks {
		if task.ID == id && task.Completed {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("task not found in tasks or not completed")
	}

	// 存在しないidを指定
	testTaskID := 100
	err = handleComplete(testTaskID)
	if err == nil {
		t.Fatalf("expected error but got nil")
	}
}

func TestHandleDelete(t *testing.T) {
	// テスト用のタスク名
	taskName := "test task"
	// テスト用のタスクを追加
	id, err := handleAdd(taskName)
	if err != nil {
		t.Fatalf("failed to add task: %v", err)
	}

	// 追加したタスクがtasks.jsonに保存されているか確認
	tasks, err := LoadTasks(tasksFileName)
	if err != nil {
		t.Fatalf("failed to load tasks: %v", err)
	}

	// 追加したタスクがtasksに含まれているか確認
	found := false
	for _, task := range tasks {
		if task.ID == id {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("task not found in tasks")
	}

	// タスクを削除する
	err = handleDelete(id)
	if err != nil {
		t.Fatalf("failed to delete task: %v", err)
	}

	// 削除したタスクがtasks.jsonに保存されているか確認
	tasks, err = LoadTasks(tasksFileName)
	if err != nil {
		t.Fatalf("failed to load tasks: %v", err)
	}

	// 削除したタスクがtasksに含まれていないか確認
	found = false
	for _, task := range tasks {
		if task.ID == id {
			found = true
			break
		}
	}
	if found {
		t.Fatalf("task found in tasks")
	}

	// 存在しないidを指定
	testTaskID := 100
	err = handleDelete(testTaskID)
	if err == nil {
		t.Fatalf("expected error but got nil")
	}
}
