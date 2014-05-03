package task

import "testing"

func TestNewTaskWithEmptyTitle(t *testing.T) {
	_, err := NewTask("")
	if err == nil {
		t.Error("Expected 'Empty Title', error, got nil")
	}
}