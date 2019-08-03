package main

import (
	"github.com/jinzhu/gorm"
)

// 数据模型
type Todo struct {
	gorm.Model
	Content     string
	IsCompleted bool
}
