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
	lastID int64
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (taskManager *TaskManager) save(task *Task) {
	if task.ID == 0 {
		taskManager.lastID++
		task.ID = taskManager.lastID
		copy := *task
		taskManager.tasks = append(taskManager.tasks, &copy)
	}
}

func (taskManager *TaskManager) GetAll() []*Task {
	return taskManager.tasks
}
