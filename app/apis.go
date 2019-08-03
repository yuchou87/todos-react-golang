package main

import (
	"github.com/gin-gonic/gin"
)

func pingHandle(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func getTodo(c *gin.Context) {
	c.JSON(200, gin.H{"message": "get"})
}

func createTodo(c *gin.Context) {
	c.JSON(200, gin.H{"message": "create"})
}

func deleteTodo(c *gin.Context) {
	c.JSON(200, gin.H{"message": "delete"})
}
