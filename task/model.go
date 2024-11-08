package task

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

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

func InitDB(dbName string) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.AutoMigrate(&Task{})
	DB = db
}

func AddTask(name string, priority Priority, due time.Time) {
	task := CreateTask(name, priority, due)
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

func MarkAsDone(id int) {
	result := DB.Model(&Task{}).Where("id = ?", id).Update("done", true)
	if result.RowsAffected == 0 {
		fmt.Println("Task not found")
	} else {
		fmt.Printf("Task %d marked as done", id)
	}
}

func LoadTasks() []Task {
	var tasks []Task
	DB.Find(&tasks)
	return tasks
}