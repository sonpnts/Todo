package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sonpnts/todo-list/handlers"
	"github.com/sonpnts/todo-list/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSConfig())
	r.POST("/tasks", handlers.CreateTask)
	r.GET("/tasks", handlers.GetTasks)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

	return r
}
