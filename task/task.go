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
	tasks []*Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (taskManager *TaskManager) save(task *Task) {
	copy := *task
	taskManager.tasks = append(taskManager.tasks, &copy)
}

func (taskManager *TaskManager) GetAll() []*Task {
	return taskManager.tasks
}
