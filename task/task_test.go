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

// func TestMarkAsDone(t *testing.T) {
// 	tasks := []Task{
// 		{ID: 1, Name: "task 1", Done: false},
// 	}
// 	MarkAsDone(&tasks, 1)

// 	if !tasks[0].Done {
// 		t.Errorf("Expected task to be mark as done")
// 	}
// }

// func TestSaveAndLoadTasks(t *testing.T) {
// 	testFile := "test_tasks.json"
// 	defer os.Remove(testFile)

// 	tasks := []Task{
// 		{ID: 1, Name: "task 1"},
// 		{ID: 2, Name: "task 2", Done: true},
// 	}

// 	SaveTasks(testFile, tasks)
// 	loadedTasks := LoadTasks(testFile)

// 	if len(loadedTasks) != 2 {
// 		t.Errorf("Expected 2 tasks, got %d", len(loadedTasks))
// 	}

// 	if loadedTasks[1].Name != "task 2" || !loadedTasks[1].Done {
// 		t.Error("Loaded task doesn't match saved task")
// 	}
// }

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