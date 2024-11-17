package services

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/Elianamos29/go-task-manager-cli/db"
	"github.com/Elianamos29/go-task-manager-cli/models"
)

func AddTask(name string, priority models.Priority, due time.Time, tags string) {
	priority = models.Priority(strings.ToLower(string(priority)))

	task := CreateTask(name, priority, due)
	task.Tags = tags
	db.DB.Create(&task)
	fmt.Printf("Added task: %s (Priority: %s, Due: %s)\n", name, priority, due.Format("2006-01-02"))
}

func DeleteTask(id int) {
	result := db.DB.Delete(&models.Task{}, id)
	if result.RowsAffected == 0 {
		fmt.Println("Task not found")
	} else {
		fmt.Printf("Task %d deleted", id)
	}
}

func MarkAsDone(id int) {
	result := db.DB.Model(&models.Task{}).Where("id = ?", id).Update("done", true)
	if result.RowsAffected == 0 {
		fmt.Println("Task not found")
	} else {
		fmt.Printf("Task %d marked as done\n", id)
	}
}

func CreateTask(name string, priority models.Priority, dueDate time.Time) models.Task {
	return models.Task{
		Name:		name,
		Done:		false,
		Priority: 	priority,
		DueDate: 	dueDate,
		Tags:		"",
	}
}

func LoadTasks() []models.Task {
	var tasks []models.Task
	db.DB.Find(&tasks)
	return tasks
}

func sortTaskByPriority(tasks *[]models.Task) {
	priorities := map[models.Priority]int{models.High: 3, models.Medium: 2, models.Low: 1}

	sort.Slice(*tasks, func(i, j int) bool {
		if priorities[(*tasks)[i].Priority] != priorities[(*tasks)[j].Priority] {
			return priorities[(*tasks)[i].Priority] > priorities[(*tasks)[j].Priority]
		}

		if (*tasks)[i].DueDate.IsZero() {
			return false
		}

		if (*tasks)[j].DueDate.IsZero() {
			return true
		}

		return (*tasks)[i].DueDate.Before((*tasks)[j].DueDate)
	})
}

func sortTaskByDueDate(tasks *[]models.Task) {
	sort.Slice(*tasks, func(i, j int) bool {
		if !(*tasks)[i].DueDate.Equal((*tasks)[j].DueDate) {
			if (*tasks)[i].DueDate.IsZero() {
				return false
			}
	
			if (*tasks)[j].DueDate.IsZero() {
				return true
			}
	
			return (*tasks)[i].DueDate.Before((*tasks)[j].DueDate)
		}

		return (*tasks)[i].Priority > (*tasks)[j].Priority
	})
}

func sortTaskByTags(tasks *[]models.Task) {
	sort.Slice(*tasks, func(i, j int) bool {
		if (*tasks)[i].Tags == "" && (*tasks)[j].Tags != "" {
			return false
		}

		if (*tasks)[i].Tags != "" && (*tasks)[j].Tags == "" {
			return true
		}

		return (*tasks)[i].Tags < (*tasks)[j].Tags
	})
}

func SortTasks(tasks *[]models.Task, sortBy string) {
	switch sortBy {
	case "priority":
		sortTaskByPriority(tasks)
	case "due":
		sortTaskByDueDate(tasks)
	case "tags":
		sortTaskByTags(tasks)
	default:
		fmt.Println("Invalid sort option! defaulting to sort by due date")
		sortTaskByDueDate(tasks)
	}
}

func FilterTasksByPriority(tasks []models.Task, priority models.Priority) []models.Task {
	priority = models.Priority(strings.ToLower(string(priority)))

	var filteredTasks []models.Task
	for _, task := range tasks {
		if task.Priority == priority {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks
}

func FilterTasksByTag(tasks []models.Task, tag string) []models.Task {
	var filteredTasks []models.Task
	for _, task := range tasks {
		tags := strings.Split(task.Tags, ",")
		for _, t := range tags {
			if strings.TrimSpace(t) == tag {
				filteredTasks = append(filteredTasks, task)
				break
			}
		}
	}

	return filteredTasks
}

func SearchTaskByName(tasks []models.Task, keyword string) []models.Task {
	keyword = strings.ToLower(keyword)

	var results []models.Task
	for _, task := range tasks {
		if strings.Contains(strings.ToLower(task.Name), keyword) {
			results = append(results, task)
		}
	}

	return results
}