package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func pingHandle(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func getTodo(c *gin.Context) {
	var todos []Todo
	if err := db.Find(&todos).Error; err != nil {
		c.JSON(502, gin.H{"message": "false", "data": ""})
		return
	}
	fmt.Println(todos)
	c.JSON(200, gin.H{"message": "true", "data": todos})
}

func createTodo(c *gin.Context) {
	var todo Todo
	c.BindJSON(&todo)
	if err := db.Create(&todo).Error; err != nil {
		c.JSON(502, gin.H{"message": "false", "data": ""})
		return
	}
	c.JSON(200, gin.H{"message": "true", "data": ""})
}

func updateTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Todo
	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		c.JSON(502, gin.H{"message": "false", "data": ""})
		return
	}
	db.Model(&todo).Update("is_completed", true)
	c.JSON(200, gin.H{"message": "true", "data": ""})
}

func deleteTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Todo
	if err := db.Where("id = ?", id).Delete(&todo).Error; err != nil {
		c.JSON(502, gin.H{"message": "false", "data": ""})
		return
	}
	c.JSON(200, gin.H{"message": "true", "data": ""})
}
