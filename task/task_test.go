package task

import (
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