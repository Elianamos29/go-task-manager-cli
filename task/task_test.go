package task

import (
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}
	db.AutoMigrate(&Task{})
	DB = db
}

func TestAddTask(t *testing.T) {
	setupTestDB()

	AddTask("Test task", High, time.Now())
	var task Task
	if err := DB.First(&task).Error; err != nil {
		t.Fatalf("Expected task to be created, got error: %v", err)
	}
}

func TestDeleteTask(t *testing.T) {
	setupTestDB()

	task := CreateTask("to be deleted", Low, time.Now())
	DB.Create(&task)
	DeleteTask(task.ID)

	var result Task
	if err := DB.First(&result, task.ID); err == nil {
		t.Fatal("Expected task to be deleted but still exists")
	}
}

func TestMarkAsDone(t *testing.T) {
	setupTestDB()

	task := CreateTask("Task done", Medium, time.Now())
	DB.Create(&task)
	MarkAsDone(task.ID)

	var result Task
	DB.First(&result, task.ID)

	if !result.Done {
		t.Errorf("Expected task to be mark as done")
	}
}

func TestLoadTasks(t *testing.T) {
	setupTestDB()

	task1 := CreateTask("Task 1", Low, time.Now())
	task2 := CreateTask("Task 2", High, time.Now())
	DB.Create(&task1)
	DB.Create(&task2)

	tasks := LoadTasks()

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
}

// func TestSortTasks(t *testing.T) {
// 	dueDate1, _ := time.Parse("2006-01-02", "2024-11-10")
// 	dueDate2, _ := time.Parse("2006-01-02", "2024-11-11")
// 	dueDate3, _ := time.Parse("2006-01-02", "2024-11-12")

// 	tasks := []Task{
//         {ID: 1, Name: "Low Priority", Priority: Low, DueDate: dueDate1},
//         {ID: 2, Name: "High Priority", Priority: High, DueDate: dueDate3},
//         {ID: 3, Name: "Medium Priority", Priority: Medium, DueDate: dueDate2},
//     }

// 	SortTasks(&tasks, "priority")

// 	if tasks[0].Priority != High || tasks[1].Priority != Medium || tasks[2].Priority != Low {
//         t.Errorf("Tasks not sorted by priority correctly")
//     }

// 	SortTasks(&tasks, "due")

// 	if tasks[0].ID != 1 || tasks[1].ID != 3 || tasks[2].ID != 2 {
//         t.Errorf("Tasks not sorted by due date correctly")
//     }
// }