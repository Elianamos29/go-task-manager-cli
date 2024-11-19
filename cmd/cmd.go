package cmd

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/Elianamos29/go-task-manager-cli/models"
	"github.com/Elianamos29/go-task-manager-cli/services"
	"github.com/Elianamos29/go-task-manager-cli/views"
)

func HandleCommands() {

	newTask := flag.String("add", "", "add a task")
	taskPriority := flag.String("priority", "medium", "Set task priority: low, medium, high")
	tags := flag.String("tags", "", "Set tags for the task(comma-separated)")
	sortBy := flag.String("sort", "due", "sort tasks by: priority, due, tags")
	deleteTaskID := flag.String("delete", "", "delete a task")
	dueDate := flag.String("due", "", "set due date for the task(YYYY-MM-DD)")
	doneTaskID := flag.String("done", "", "mark task as done")
	priority := flag.String("filter", "", "filter tasks by priority")
	keyword := flag.String("search", "", "search tasks by name")
	tagFilter := flag.String("filter-tag", "", "filter tasks by specific tag")
	showCompleted := flag.Bool("completed", false, "show completed tasks")
	showIncomplete := flag.Bool("incomplete", false, "show incomplete tasks")
	flag.Parse()

	if *newTask != "" {
		handleAddTask(*newTask, *taskPriority, *dueDate, *tags)
	}

	if *deleteTaskID != "" {
		handleDeleteTask(*deleteTaskID)
	}

	if *doneTaskID != "" {
		handleMarkAsDone(*doneTaskID)
	}

	handleDisplayTasks(*sortBy, *priority, *keyword, *tagFilter, *showCompleted, *showIncomplete)
}

func handleAddTask(name, priority, dueDate, tags string) {
	var due time.Time
	if dueDate != "" {
		parsedDate, err := parseDate(dueDate)
		if err != nil {
			fmt.Println("Invalid due date format! Please use YYYY-MM-DD")
		} else {
			due = parsedDate
		}
	}

	services.AddTask(name, models.Priority(priority), due, tags)
}

func handleDeleteTask(idStr string) {
	id, err := parseID(idStr)
	if err != nil {
		fmt.Println("invalid task ID:", id)
	} else {
		services.DeleteTask(id)
	}
}

func handleMarkAsDone(idStr string) {
	id, err := parseID(idStr)
	if err != nil {
		fmt.Println("invalid task ID:", id)
	} else {
		services.MarkAsDone(id)
	}
}

func handleDisplayTasks(sortBy, priority, keyword, tagFilter string, showCompleted, showIncomplete bool) {
	views.Display(sortBy, priority, keyword, tagFilter, showCompleted, showIncomplete)
}

func parseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

func parseID(idStr string) (int, error) {
	return strconv.Atoi(idStr)
}