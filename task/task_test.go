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

func TestSaveOneTaskAndGetOne(t *testing.T) {
	task, err := NewTask("My task")
	if err != nil {
		t.Errorf("New task : %v", err)
	}

	taskManager := NewTaskManager()
	taskManager.save(task)

	allTask := taskManager.GetAll()
	if len(allTask) != 1 {
		t.Errorf("Expected 1 task, got %v", len(allTask))
	}

	if *allTask[0] != *task {
		t.Errorf("Expected %v, got %v", task, allTask)
	}
}


