package task

import (
	"testing"
	"time"

	"github.com/Elianamos29/go-task-manager-cli/models"
	"github.com/Elianamos29/go-task-manager-cli/db"
	"github.com/Elianamos29/go-task-manager-cli/services"
	
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() {
	testdb, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}
	testdb.AutoMigrate(&models.Task{})
	db.DB = testdb
}

func TestAddTask(t *testing.T) {
	setupTestDB()

	services.AddTask("Test task", models.High, time.Now(),"test")
	var task models.Task
	if err := db.DB.First(&task).Error; err != nil {
		t.Fatalf("Expected task to be created, got error: %v", err)
	}
}

func TestDeleteTask(t *testing.T) {
	setupTestDB()

	task := services.CreateTask("to be deleted", models.Low, time.Now())
	db.DB.Create(&task)
	services.DeleteTask(task.ID)

	var result models.Task
	if err := db.DB.First(&result, task.ID); err == nil {
		t.Fatal("Expected task to be deleted but still exists")
	}
}

func TestMarkAsDone(t *testing.T) {
	setupTestDB()

	task := services.CreateTask("Task done", models.Medium, time.Now())
	db.DB.Create(&task)
	services.MarkAsDone(task.ID)

	var result models.Task
	db.DB.First(&result, task.ID)

	if !result.Done {
		t.Errorf("Expected task to be mark as done")
	}
}

func TestLoadTasks(t *testing.T) {
	setupTestDB()

	task1 := services.CreateTask("Task 1", models.Low, time.Now())
	task2 := services.CreateTask("Task 2", models.High, time.Now())
	db.DB.Create(&task1)
	db.DB.Create(&task2)

	tasks := services.LoadTasks()

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
}

func TestSortTasks(t *testing.T) {
	dueDate1, _ := time.Parse("2006-01-02", "2024-11-10")
	dueDate2, _ := time.Parse("2006-01-02", "2024-11-11")
	dueDate3, _ := time.Parse("2006-01-02", "2024-11-12")

	tasks := []models.Task{
        {ID: 1, Name: "Low Priority", Priority: models.Low, DueDate: dueDate1},
        {ID: 2, Name: "High Priority", Priority: models.High, DueDate: dueDate3},
        {ID: 3, Name: "Medium Priority", Priority: models.Medium, DueDate: dueDate2},
    }

	services.SortTasks(&tasks, "priority")

	if tasks[0].Priority != models.High || tasks[1].Priority != models.Medium || tasks[2].Priority != models.Low {
        t.Errorf("Tasks not sorted by priority correctly")
    }

	services.SortTasks(&tasks, "due")

	if tasks[0].ID != 1 || tasks[1].ID != 3 || tasks[2].ID != 2 {
        t.Errorf("Tasks not sorted by due date correctly")
    }
}