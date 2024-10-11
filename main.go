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
	deleteTaskID := flag.String("delete", "", "delete a task")
	doneTaskID := flag.String("done", "", "mark task as done")
	showCompleted := flag.Bool("completed", false, "show completed tasks")
	showIncomplete := flag.Bool("incomplete", false, "show incomplete tasks")
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

	if *deleteTaskID != "" {
		id, err := strconv.Atoi(*deleteTaskID)
		if err == nil {
			deleteTask(&tasks, id)
			saveTasks(tasks)
		} else {
			fmt.Println("invalid task id")
		}
	}

	fmt.Println("Your tasks:")
	if *showCompleted && *showIncomplete {
		fmt.Println("Please specify only one filter: --completed or --incomplete.")
	} else if *showCompleted {
		displayTasks(tasks, &[]bool{true}[0])
	} else if *showIncomplete {
		displayTasks(tasks, &[]bool{false}[0])
	} else {
		displayTasks(tasks, nil)
	}
}

func addTask(tasks *[]Task, name string) {
	newID := len(*tasks) + 1
	*tasks = append(*tasks, Task{ID: newID, Name: name, Done: false})
	fmt.Printf("Added task: %s\n", name)
}

func deleteTask(tasks *[]Task, id int) {
	for i, task := range *tasks {
		if task.ID == id {
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			fmt.Printf("Task %d deleted\n", id)
			return
		}
	}

	fmt.Println("Task not found")
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
	file, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Error marshaling tasks:", err)
		return
	}
	err = os.WriteFile(taskFile, file, 0644)
	if err != nil {
		fmt.Println("Error writing to a file:", err)
		return
	}
}

func loadTasks() []Task {
	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
		return []Task{}
	}

	file, err := os.ReadFile(taskFile)
	if err != nil {
		fmt.Println("Error reading from a file:", err)
		return []Task{}
	}

	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		fmt.Println("Error marshaling tasks:", err)
		return []Task{}
	}

	return tasks
}

func displayTasks(tasks []Task, filter *bool) {
	for _, task := range tasks {
		if filter != nil {
			if *filter && !task.Done {
				continue
			}
			if !*filter && task.Done {
				continue
			}
		}
		status := "Not done"
		if task.Done {
			status = "Done"
		}

		fmt.Printf("%d. %s [%s]\n", task.ID, task.Name, status)
	}
}