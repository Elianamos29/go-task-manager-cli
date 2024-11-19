package views

import (
	"fmt"
	"strings"

	"github.com/Elianamos29/go-task-manager-cli/models"
	"github.com/Elianamos29/go-task-manager-cli/services"
)

func displayTasks(tasks []models.Task, filter *bool) {
	fmt.Println("Your tasks")
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

		due := "No due date"
		if !task.DueDate.IsZero() {
			due = task.DueDate.Format("2006-01-02")
		}

		tags := "No tags"
		if task.Tags != "" {
			tags = task.Tags
		}

		fmt.Printf("%d. %s [%s] (Priority: %s, Due: %s, Tags: %s)\n", task.ID, task.Name, status, task.Priority, due, tags)
	}
}

func Display(sortBy, priority, keyword, tagFilter string, showCompleted, showIncomplete bool) {
	tasks := services.LoadTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks to display.")
		return
	}

	if priority != "" {
		tasks = services.FilterTasksByPriority(tasks, models.Priority(priority))
	}

	if keyword != "" {
		tasks = services.SearchTaskByName(tasks, strings.TrimSpace(keyword))
	}

	if tagFilter != "" {
		tasks = services.FilterTasksByTag(tasks, strings.TrimSpace(tagFilter))
	}

	services.SortTasks(&tasks, sortBy)
	if showCompleted && showIncomplete {
		fmt.Println("Please specify only one filter: --completed or --incomplete.")
	} else if showCompleted {
		displayTasks(tasks, &[]bool{true}[0])
	} else if showIncomplete {
		displayTasks(tasks, &[]bool{false}[0])
	} else {
		displayTasks(tasks, nil)
	}
}