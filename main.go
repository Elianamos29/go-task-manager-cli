package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var taskFile = "tasks.json"

func main() {
	tasks := loadTasks()

	newTask := flag.String("add", "", "add a task")
	doneTaskID := flag.String("done", "", "mark task as done")
	flag.Parse()

	if *newTask != "" {
		addTask(&tasks, *newTask)
		saveTasks(tasks)
	}

	if *doneTaskID != "" {
		id, err := strconv.Atoi(*doneTaskID)
		if err == nil {
			markAsDone(&tasks, id)
			saveTasks(tasks)
		} else {
			fmt.Println("Invalid task id")
		}
	}

	fmt.Println("Your tasks:")
	for _, task := range tasks {
		status := "Not done"
		if task.Done {
			status = "Done"
		}

		fmt.Printf("%d. %s [%s]\n", task.ID, task.Name, status)
	}
}

func addTask(tasks *[]Task, name string) {
	newID := len(*tasks) + 1
	*tasks = append(*tasks, Task{ID: newID, Name: name, Done: false})
	fmt.Printf("Added task: %s\n", name)
}

func markAsDone(tasks *[]Task, id int) {
	for i, task := range *tasks {
		if task.ID == id {
			(*tasks)[i].Done = true
			fmt.Printf("Task %d marked as done\n", id)
			return
		}
	}

	fmt.Println("Task not found")
}

func saveTasks(tasks []Task) {
	file, _ := json.MarshalIndent(tasks, "", " ")
	_ = os.WriteFile(taskFile, file, 0644)
}

func loadTasks() []Task {
	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
		return []Task{}
	}

	file, _ := os.ReadFile(taskFile)
	var tasks []Task
	_ = json.Unmarshal(file, &tasks)
	return tasks
}