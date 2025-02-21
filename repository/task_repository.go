package repository

import (
	"errors"
	"sync"

	"github.com/sonpnts/todo-list/models"
)

var (
	tasks  = make(map[int]models.Task)
	lastID = 0
	mu     sync.Mutex
)

func CreateTask(task models.Task) int {
	mu.Lock()
	defer mu.Unlock()
	lastID++
	task.ID = lastID
	tasks[task.ID] = task
	return task.ID
}

func GetTasks() []models.Task {
	mu.Lock()
	defer mu.Unlock()
	taskList := make([]models.Task, 0, len(tasks))
	for _, task := range tasks {
		taskList = append(taskList, task)
	}
	return taskList
}

func UpdateTask(id int, completed bool) (models.Task, error) {
	mu.Lock()
	defer mu.Unlock()
	task, exists := tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}
	task.Completed = completed
	tasks[id] = task
	return task, nil
}

func DeleteTask(id int) error {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(tasks, id)
	return nil
}
