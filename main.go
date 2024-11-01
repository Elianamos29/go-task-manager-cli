package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/Elianamos29/go-task-manager-cli/task"
)

var taskFile = "tasks.json"

func main() {
	tasks := task.LoadTasks(taskFile)

	newTask := flag.String("add", "", "add a task")
	taskPriority := flag.String("priority", "medium", "Set task priority: low, medium, high")
	sortBy := flag.String("sort", "due", "sort tasks by: priority, due")
	deleteTaskID := flag.String("delete", "", "delete a task")
	dueDate := flag.String("due", "", "set due date for the task(YYYY-MM-DD)")
	doneTaskID := flag.String("done", "", "mark task as done")
	showCompleted := flag.Bool("completed", false, "show completed tasks")
	showIncomplete := flag.Bool("incomplete", false, "show incomplete tasks")
	flag.Parse()

	if *newTask != "" {
		task.AddTask(&tasks, *newTask, *taskPriority, *dueDate)
		task.SaveTasks(taskFile, tasks)
	}

	if *doneTaskID != "" {
		id, err := strconv.Atoi(*doneTaskID)
		if err != nil {
			fmt.Println("invalid task ID:", *doneTaskID)
		} else {
			task.MarkAsDone(&tasks, id)
			task.SaveTasks(taskFile, tasks)
		}
	}

	if *deleteTaskID != "" {
		id, err := strconv.Atoi(*deleteTaskID)
		if err != nil {
			fmt.Println("invalid task ID:", *deleteTaskID)
		} else {
			task.DeleteTask(&tasks, id)
			task.SaveTasks(taskFile, tasks)
		}
	}

	fmt.Println("Your tasks:")
	task.SortTasks(&tasks, *sortBy)
	if *showCompleted && *showIncomplete {
		fmt.Println("Please specify only one filter: --completed or --incomplete.")
	} else if *showCompleted {
		task.DisplayTasks(tasks, &[]bool{true}[0])
	} else if *showIncomplete {
		task.DisplayTasks(tasks, &[]bool{false}[0])
	} else {
		task.DisplayTasks(tasks, nil)
	}
}