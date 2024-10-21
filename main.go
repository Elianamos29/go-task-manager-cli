package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var taskFile = "tasks.json"
var maxID int

func main() {
	tasks := loadTasks()

	newTask := flag.String("add", "", "add a task")
	taskPriority := flag.String("priority", "medium", "Set task priority: low, medium, high")
	deleteTaskID := flag.String("delete", "", "delete a task")
	doneTaskID := flag.String("done", "", "mark task as done")
	showCompleted := flag.Bool("completed", false, "show completed tasks")
	showIncomplete := flag.Bool("incomplete", false, "show incomplete tasks")
	flag.Parse()

	if *newTask != "" {
		addTask(&tasks, *newTask, *taskPriority)
		saveTasks(tasks)
	}

	if *doneTaskID != "" {
		id, err := strconv.Atoi(*doneTaskID)
		if err != nil {
			fmt.Println("invalid task ID:", *doneTaskID)
		} else {
			markAsDone(&tasks, id)
			saveTasks(tasks)
		}
	}

	if *deleteTaskID != "" {
		id, err := strconv.Atoi(*deleteTaskID)
		if err != nil {
			fmt.Println("invalid task ID:", *deleteTaskID)
		} else {
			deleteTask(&tasks, id)
			saveTasks(tasks)
		}
	}

	fmt.Println("Your tasks:")
	sortTaskByPriority(&tasks)
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

func addTask(tasks *[]Task, name string, priority string) {
	maxID++

	var taskPriority Priority

	switch priority {
	case "low":
		taskPriority = Low
	case "medium":
		taskPriority = Medium
	case "high":
		taskPriority = High
	default:
		fmt.Println("Invalid priority! defaulting to 'medium'")
		taskPriority = Medium
	}

	*tasks = append(*tasks, Task{ID: maxID, Name: name, Done: false, Priority: taskPriority})
	fmt.Printf("Added task: %s (Priority: %s)\n", name, taskPriority)
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

	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	return tasks
}

func sortTaskByPriority(tasks *[]Task) {
	sort.Slice(*tasks, func(i, j int) bool {
		priorities := map[Priority]int{High: 3, Medium: 2, Low: 1}
		return priorities[(*tasks)[i].Priority] > priorities[(*tasks)[j].Priority]
	})
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

		fmt.Printf("%d. %s [%s] (Priority: %s)\n", task.ID, task.Name, status, task.Priority)
	}
}