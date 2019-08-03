package main

import (
	"github.com/gin-gonic/gin"
)

func todosFunc(r *gin.RouterGroup) {
	todo := r.Group("/todos")
	{
		todo.GET("/", getTodo)
		todo.POST("/", createTodo)
		todo.PUT("/:id", updateTodo)
		todo.DELETE("/:id", deleteTodo)
	}
}

func applyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		v1.GET("/ping", pingHandle)
		todosFunc(v1)
	}
}

func setRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		applyRoutes(api)
	}
}
