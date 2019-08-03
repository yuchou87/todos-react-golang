package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pingHandle(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
