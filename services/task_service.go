package services

import (
	"github.com/sonpnts/todo-list/models"
	"github.com/sonpnts/todo-list/repository"
)

func CreateTask(task models.Task) int {
	return repository.CreateTask(task)
}

func GetTasks(page, pageSize int) ([]models.Task, bool, error) {
	tasks, hasNextPage := repository.GetTasks(page, pageSize)
	return tasks, hasNextPage, nil
}

func SearchTasks(query string, page, pageSize int) ([]models.Task, bool, error) {
	tasks, hasNextPage := repository.SearchTasks(query, page, pageSize)
	return tasks, hasNextPage, nil
}

func UpdateTask(id int, completed bool) (models.Task, error) {
	return repository.UpdateTask(id, completed)
}

func DeleteTask(id int) error {
	return repository.DeleteTask(id)
}
