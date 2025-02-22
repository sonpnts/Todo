package handlers

import (
	"github.com/sonpnts/todo-list/models"
	"github.com/sonpnts/todo-list/repository"
	"github.com/sonpnts/todo-list/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	id := services.CreateTask(newTask)
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func GetTasks(c *gin.Context) {
	// Lấy tham số `page` từ query, mặc định là 1
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	// Lấy tham số `pageSize`, mặc định là 10
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// Gọi repository để lấy dữ liệu
	tasks, hasNextPage := repository.GetTasks(page, pageSize)

	var nextPage *int
	if hasNextPage {
		next := page + 1
		nextPage = &next
	}
	// Trả về kết quả
	c.JSON(http.StatusOK, gin.H{
		"page":     page,
		"pageSize": pageSize,
		"tasks":    tasks,
		"nextPage": nextPage,
	})
}
func UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updateData struct {
		Completed bool `json:"completed"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	task, err := services.UpdateTask(id, updateData.Completed)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = services.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
