package task

import "fmt"

type Task struct {
	ID    int64
	Title string
	Done  bool
}

func NewTask(title string) (*Task, error) {
	if title == "" {
		return nil, fmt.Errorf("Empty title")
	}
	return &Task{0, title, false}, nil
}


type TaskManager struct {
	task *Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (taskManager *TaskManager) save(task *Task) {
	taskManager.task = task
}

func (taskManager *TaskManager) GetAll() []*Task {
	return []*Task{taskManager.task}
}
