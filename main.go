package main

import (
	"flag"
	"fmt"
	"strconv"
)

var tasks = []Task{
	{ID: 1, Name: "Learn Go", Done: false},
	{ID: 2, Name: "Write cli", Done: false},
}

func main() {
	newTask := flag.String("add", "", "add a task")
	doneTaskID := flag.String("done", "", "mark task as done")
	flag.Parse()

	if *newTask != "" {
		addTask(*newTask)
	}

	if *doneTaskID != "" {
		id, err := strconv.Atoi(*doneTaskID)
		if err == nil {
			markAsDone(id)
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

func addTask(name string) {
	newID := len(tasks) + 1
	tasks = append(tasks, Task{ID: newID, Name: name, Done: false})
	fmt.Printf("Added task: %s\n", name)
}

func markAsDone(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			fmt.Printf("Task %d marked as done\n", id)
			return
		}
	}

	fmt.Println("Task not found")
}