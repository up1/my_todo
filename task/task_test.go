package task

import "testing"

func TestNewTaskWithEmptyTitle(t *testing.T) {
	_, err := NewTask("")
	if err == nil {
		t.Error("Expected 'Empty Title', error, got nil")
	}
}

func TestNewTask(t *testing.T) {
	title := "My Title"
	task, _ := NewTask(title)
	if task.Title != title {
		t.Errorf("Expected title %q, got %q", title, task.Title)
	}
	if task.Done {
		t.Errorf("New task is done")
	}
}


