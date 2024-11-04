package task

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"sort"
	"time"
)

var DB *gorm.DB

var maxID int
type Priority string

const (
	Low Priority = "low"
	Medium Priority = "medium"
	High Priority = "high"
)

type Task struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
	Done bool `json:"done"`
	Priority Priority `json:"priority"`
	DueDate time.Time `json:"due_date"`
}

func initDB(dbName string) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.AutoMigrate(&Task{})
	DB = db
}

func AddTask(tasks *[]Task, name string, priority Priority, due time.Time) {
	task := Task{Name: name, Done: false, Priority: priority, DueDate: due}
	DB.Create(&task)
	fmt.Printf("Added task: %s (Priority: %s, Due: %s)\n", name, priority, due.Format("2006-01-02"))
}

func DeleteTask(id int) {
	result := DB.Delete(&Task{}, id)
	if result.RowsAffected == 0 {
		fmt.Println("Task not found")
	} else {
		fmt.Printf("Task %d deleted", id)
	}
}

func MarkAsDone(tasks *[]Task, id int) {
	for i, task := range *tasks {
		if task.ID == id {
			(*tasks)[i].Done = true
			fmt.Printf("Task %d marked as done\n", id)
			return
		}
	}

	fmt.Println("Task not found")
}

func SaveTasks(taskFile string, tasks []Task) {
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

func LoadTasks(taskFile string) []Task {
	var tasks []Task
	DB.Find(&tasks)
	return tasks
}

func sortTaskByPriority(tasks *[]Task) {
	priorities := map[Priority]int{High: 3, Medium: 2, Low: 1}

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

func sortTaskByDueDate(tasks *[]Task) {
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

func SortTasks(tasks *[]Task, sortBy string) {
	switch sortBy {
	case "priority":
		sortTaskByPriority(tasks)
	case "due":
		sortTaskByDueDate(tasks)
	default:
		fmt.Println("Invalid sort option! defaulting to sort by due date")
		sortTaskByDueDate(tasks)
	}
}

func DisplayTasks(tasks []Task, filter *bool) {
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

		fmt.Printf("%d. %s [%s] (Priority: %s, Due: %s)\n", task.ID, task.Name, status, task.Priority, due)
	}
}