package task

import (
	"os"
	"testing"
	"time"
)

func TestAddTask(t *testing.T) {
	tasks := []Task{}
	title := "Test Task"
	dueDate, _ := time.Parse("2006-01-02", "2024-11-12")

	AddTask(&tasks, title, string(High), dueDate.Format("2006-01-02"))

	if len(tasks) != 1 {
		t.Errorf("Expected %d task, got %d", 1, len(tasks))
	}

	if tasks[0].Name != title {
		t.Errorf("expected task name %q, got %q", title, tasks[0].Name)
	}

	if tasks[0].Priority != High {
		t.Errorf("expected task priority %q, got %q", High, tasks[0].Priority)
	}

	if !tasks[0].DueDate.Equal(dueDate) {
		t.Errorf("Expected task due date %q, got %q", dueDate.Format("2006-01-02"), tasks[0].DueDate.Format("2006-01-02"))
	}
}

func TestDeleteTask(t *testing.T) {
	tasks := []Task{
		{ID: 1, Name: "task 1"},
		{ID: 2, Name: "task 2"},
	}

	DeleteTask(&tasks, 1)

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task after deletion, got %d", len(tasks))
	}

	if tasks[0].ID != 2 {
		t.Errorf("Expected remaining task with ID 2, got %d", tasks[0].ID)
	}
}

func TestMarkAsDone(t *testing.T) {
	tasks := []Task{
		{ID: 1, Name: "task 1", Done: false},
	}
	MarkAsDone(&tasks, 1)

	if !tasks[0].Done {
		t.Errorf("Expected task to be mark as done")
	}
}

func TestSaveAndLoadTasks(t *testing.T) {
	testFile := "test_tasks.json"
	defer os.Remove(testFile)

	tasks := []Task{
		{ID: 1, Name: "task 1"},
		{ID: 2, Name: "task 2", Done: true},
	}

	SaveTasks(testFile, tasks)
	loadedTasks := LoadTasks(testFile)

	if len(loadedTasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(loadedTasks))
	}

	if loadedTasks[1].Name != "task 2" || !loadedTasks[1].Done {
		t.Error("Loaded task doesn't match saved task")
	}
}

func TestSortTasks(t *testing.T) {
	dueDate1, _ := time.Parse("2006-01-02", "2024-11-10")
	dueDate2, _ := time.Parse("2006-01-02", "2024-11-11")
	dueDate3, _ := time.Parse("2006-01-02", "2024-11-12")

	tasks := []Task{
        {ID: 1, Name: "Low Priority", Priority: Low, DueDate: dueDate1},
        {ID: 2, Name: "High Priority", Priority: High, DueDate: dueDate3},
        {ID: 3, Name: "Medium Priority", Priority: Medium, DueDate: dueDate2},
    }

	SortTasks(&tasks, "priority")

	if tasks[0].Priority != High || tasks[1].Priority != Medium || tasks[2].Priority != Low {
        t.Errorf("Tasks not sorted by priority correctly")
    }

	SortTasks(&tasks, "due")

	if tasks[0].ID != 1 || tasks[1].ID != 3 || tasks[2].ID != 2 {
        t.Errorf("Tasks not sorted by due date correctly")
    }
}