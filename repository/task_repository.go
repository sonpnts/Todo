package repository

import (
	"errors"
	"sort"
	"strings"
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

func GetTasks(page, pageSize int) ([]models.Task, bool) {
	mu.Lock()
	defer mu.Unlock()

	taskList := make([]models.Task, 0, len(tasks))
	for _, task := range tasks {
		taskList = append(taskList, task)
	}

	sort.Slice(taskList, func(i, j int) bool {
		return taskList[i].ID < taskList[j].ID
	})

	start := (page - 1) * pageSize
	end := start + pageSize

	if start >= len(taskList) {
		return []models.Task{}, false
	}

	if end > len(taskList) {
		end = len(taskList)
	}

	hasNextPage := end < len(taskList)

	return taskList[start:end], hasNextPage
}

func SearchTasks(query string, page, pageSize int) ([]models.Task, bool) {
	mu.Lock()
	defer mu.Unlock()

	var filteredTasks []models.Task
	lowerQuery := strings.ToLower(query)

	for _, task := range tasks {
		if strings.Contains(strings.ToLower(task.Title), lowerQuery) {
			filteredTasks = append(filteredTasks, task)
		}
	}

	sort.Slice(filteredTasks, func(i, j int) bool {
		return filteredTasks[i].ID < filteredTasks[j].ID
	})

	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize

	if startIndex >= len(filteredTasks) {
		return []models.Task{}, false
	}

	if endIndex > len(filteredTasks) {
		endIndex = len(filteredTasks)
	}

	hasNextPage := endIndex < len(filteredTasks)
	return filteredTasks[startIndex:endIndex], hasNextPage
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
