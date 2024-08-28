package main

import (
	"encoding/json"
	"log"
	"os"
)

const tasksFileName = "tasks.json"

func SaveTasks(fileName string, tasks []*Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func LoadTasks(fileName string) ([]*Task, error) {
	if !fileExists(fileName) {
		log.Println("File does not exist, returning empty task list.")
		return []*Task{}, nil
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var tasks []*Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}
