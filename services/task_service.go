package services

import (
	"github.com/sonpnts/todo-list/models"
	"github.com/sonpnts/todo-list/repository"
)

func CreateTask(task models.Task) int {
	return repository.CreateTask(task)
}

func GetTasks() []models.Task {
	return repository.GetTasks()
}

func UpdateTask(id int, completed bool) (models.Task, error) {
	return repository.UpdateTask(id, completed)
}

func DeleteTask(id int) error {
	return repository.DeleteTask(id)
}
