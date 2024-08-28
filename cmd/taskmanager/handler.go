package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func handleArgs() {
	if len(os.Args) < 2 {
		log.Fatal("No command provided")
	}
	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			log.Fatal("No task provided")
		}
		taskName := os.Args[2]
		_, err := handleAdd(taskName)
		if err != nil {
			log.Fatalf("Error adding task: %v", err)
		}
	case "list":
		err := handleList()
		if err != nil {
			log.Fatalf("Error listing tasks: %v", err)
		}
	case "complete":
		if len(os.Args) < 3 {
			log.Fatal("No task ID provided")
		}
		taskIDStr := os.Args[2]
		taskID, err := strconv.Atoi(taskIDStr)
		if err != nil {
			log.Fatalf("Invalid task ID: %v", err)
		}
		err = handleComplete(taskID)
		if err != nil {
			log.Fatalf("Error completing task: %v", err)
		}
	case "delete":
		if len(os.Args) < 3 {
			log.Fatal("No task ID provided")
		}
		taskIDStr := os.Args[2]
		taskID, err := strconv.Atoi(taskIDStr)
		if err != nil {
			log.Fatalf("Invalid task ID: %v", err)
		}
		err = handleDelete(taskID)
		if err != nil {
			log.Fatalf("Error deleting task: %v", err)
		}
	case "help":
		handleHelp()
	default:
		log.Fatalf("Unknown command: %s", os.Args[1])
	}
}

func handleAdd(taskName string) (int, error) {
	tasks, err := LoadTasks(tasksFileName)
	if err != nil {
		return 0, err
	}
	newTaskID := getNextTaskID(tasks)
	newTask := NewTask(newTaskID, taskName)
	tasks = append(tasks, newTask)

	if err := SaveTasks(tasksFileName, tasks); err != nil {
		return 0, err
	}
	fmt.Printf("Task %q added successfully.\n", taskName)
	return newTaskID, nil
}

func handleList() error {
	tasks, err := LoadTasks(tasksFileName)
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks.")
		return nil
	}

	fmt.Println("Tasks:")
	for _, task := range tasks {
		var checked string
		if task.Completed {
			checked = "✓"
		} else {
			checked = " "
		}
		fmt.Printf("%s %d: %s (Created at: %s)\n", checked, task.ID, task.Name, task.CreatedAt.Format("2006-01-02 15:04:05"))
	}
	return nil
}

func handleComplete(id int) error {
	tasks, err := LoadTasks(tasksFileName)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		if task.ID == id {
			task.ToggleCompleted()
			if err := SaveTasks(tasksFileName, tasks); err != nil {
				return err
			}
			fmt.Printf("Task %d completed.\n", task.ID)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func handleDelete(id int) error {
	tasks, err := LoadTasks(tasksFileName)
	if err != nil {
		return err
	}
	for ind, task := range tasks {
		if task.ID == id {
			// idのタスクを除いたスライスを作成
			tasks = append(tasks[:ind], tasks[ind+1:]...)
			if err := SaveTasks(tasksFileName, tasks); err != nil {
				return err
			}
			fmt.Printf("Task %d deleted.\n", task.ID)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func handleHelp() {
	fmt.Println(`Usage:
		task add <task name>    Add a new task
		task list               List tasks
		task complete <task ID> Complete a task
		task delete <task ID>   Delete a task`)
}
