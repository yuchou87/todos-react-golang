package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	setRoutes(r)
	r.Run(":8080")
}
