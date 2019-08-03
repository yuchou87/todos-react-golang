package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("mssql", "sqlserver://username:password@localhost:1433?database=dbname")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	r := gin.Default()
	setRoutes(r)
	r.Run(":8080")
}
