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

func TestSaveTwoAndGetTwo(t *testing.T) {
	task1, err := NewTask("My task 1")
	if err != nil {
		t.Errorf("New task : %v", err)
	}

	task2, err := NewTask("My task 2")
	if err != nil {
		t.Errorf("New task : %v", err)
	}

	taskManager := NewTaskManager()
	taskManager.save(task1)
	taskManager.save(task2)

	allTask := taskManager.GetAll()
	if len(allTask) != 2 {
		t.Errorf("Expected 2 task, got %v", len(allTask))
	}
}

func TestSaveUpdateWithDoneAndGet(t *testing.T) {
	task, err := NewTask("My task")
	if err != nil {
		t.Errorf("New task : %v", err)
	}

	taskManager := NewTaskManager()
	taskManager.save(task)

	task.Done = true

	allTask := taskManager.GetAll()
	if allTask[0].Done {
		t.Errorf("Save task not done")
	}
}

func TestDuplicateSaveAndGetOne(t *testing.T) {
	task, err := NewTask("My task")
	if err != nil {
		t.Errorf("New task : %v", err)
	}

	taskManager := NewTaskManager()
	taskManager.save(task)
	taskManager.save(task)

	allTask := taskManager.GetAll()
	if len(allTask) != 1 {
		t.Errorf("Expected 1 task, got %v", len(allTask))
	}
}

func TestSaveAndFindByID(t *testing.T) {
	task, err := NewTask("My task")
	if err != nil {
		t.Errorf("New task : %v", err)
	}

	taskManager := NewTaskManager()
	taskManager.save(task)

	myTask, ok := taskManager.Find(task.ID)
	if !ok {
		t.Errorf("Task not found")
	}

	if *task != *myTask {
		t.Errorf("Expected %v, got %v", task, myTask)
	}
}


func TestSaveFindAndEdit(t *testing.T) {
	task, err := NewTask("My task")
	if err != nil {
		t.Errorf("New task : %v", err)
	}

	taskManager := NewTaskManager()
	taskManager.save(task)

	myTask, _ := taskManager.Find(task.ID)
	myTask.Done = true
	taskManager.save(myTask)

	allTask := taskManager.GetAll()
	if !allTask[0].Done {
		t.Errorf("Save task not done")
	}

	if len(allTask) != 1 {
		t.Errorf("Expected 1 task, got %v", len(allTask))
	}
}