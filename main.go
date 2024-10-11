package main

import (
	"fmt"
	"flag"
)

var tasks = []Task{
	{ID: 1, Name: "Learn Go", Done: false},
	{ID: 2, Name: "Write cli", Done: false},
}

func main() {
	newTask := flag.String("add", "", "add a task")
	flag.Parse()

	if *newTask != "" {
		addTask(*newTask)
	}

	fmt.Println("Your tasks:")
	for _, task := range tasks {
		fmt.Printf("%d. %s [Done: %t]\n", task.ID, task.Name, task.Done)
	}
}

func addTask(name string) {
	newID := len(tasks) + 1
	tasks = append(tasks, Task{ID: newID, Name: name, Done: false})
	fmt.Printf("Added task: %s\n", name)
}