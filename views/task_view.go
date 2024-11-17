package views

import (
	"fmt"

	"github.com/Elianamos29/go-task-manager-cli/models"
)

func DisplayTasks(tasks []models.Task, filter *bool) {
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