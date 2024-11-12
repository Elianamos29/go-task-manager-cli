package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/Elianamos29/go-task-manager-cli/db"
	"github.com/Elianamos29/go-task-manager-cli/models"
	"github.com/Elianamos29/go-task-manager-cli/services"
	"github.com/Elianamos29/go-task-manager-cli/views"
)

func main() {
	db.InitDB("tasks.db")

	newTask := flag.String("add", "", "add a task")
	taskPriority := flag.String("priority", "medium", "Set task priority: low, medium, high")
	sortBy := flag.String("sort", "due", "sort tasks by: priority, due")
	deleteTaskID := flag.String("delete", "", "delete a task")
	dueDate := flag.String("due", "", "set due date for the task(YYYY-MM-DD)")
	doneTaskID := flag.String("done", "", "mark task as done")
	priority := flag.String("filter", "", "filter tasks by priority")
	keyword := flag.String("search", "", "search tasks by name")
	showCompleted := flag.Bool("completed", false, "show completed tasks")
	showIncomplete := flag.Bool("incomplete", false, "show incomplete tasks")
	flag.Parse()

	if *newTask != "" {
		var due time.Time
		if *dueDate != "" {
			parsedDue, err := time.Parse("2006-01-02", *dueDate)
			if err != nil {
				fmt.Println("Invalid due date format! Please use YYYY-MM-DD")
			} else {
				due = parsedDue
			}
		}
		services.AddTask(*newTask, models.Priority(*taskPriority), due)
	}

	if *doneTaskID != "" {
		id, err := strconv.Atoi(*doneTaskID)
		if err != nil {
			fmt.Println("invalid task ID:", *doneTaskID)
		} else {
			services.MarkAsDone(id)
		}
	}

	if *deleteTaskID != "" {
		id, err := strconv.Atoi(*deleteTaskID)
		if err != nil {
			fmt.Println("invalid task ID:", *deleteTaskID)
		} else {
			services.DeleteTask(id)
		}
	}

	fmt.Println("Your tasks:")
	tasks := services.LoadTasks()
	if *priority != "" {
		tasks = services.FilterTasksByPriority(tasks, models.Priority(*priority))
	}

	if *keyword != "" {
		tasks = services.SearchTaskByName(tasks, *keyword)
	}

	services.SortTasks(&tasks, *sortBy)
	if *showCompleted && *showIncomplete {
		fmt.Println("Please specify only one filter: --completed or --incomplete.")
	} else if *showCompleted {
		views.DisplayTasks(tasks, &[]bool{true}[0])
	} else if *showIncomplete {
		views.DisplayTasks(tasks, &[]bool{false}[0])
	} else {
		views.DisplayTasks(tasks, nil)
	}
}