package main

import "time"

// 数据模型
type Todo struct {
	ID          uint      `gorm:"primary_key;AUTO_INCREMENT;not null" json:"id"`
	Content     string    `gorm:"type:text" json:"content" binding:"required"`
	IsCompleted bool      `gorm:"default:false" json:"is_completed"`
	CreatedAt   time.Time `gorm:"default:null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:null" json:"updated_at"`
	DeletedAt   time.Time `gorm:"default:null" json:"deleted_at"`
}
