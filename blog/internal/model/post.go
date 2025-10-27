package model

import (
	"gorm.io/gorm"
)

// 文章db模型
type Post struct {
	gorm.Model
	ID      uint
	Title   string `gorm:"type:varchar(256); not null; index"`
	Content string `gorm:"type:text; not null"`
	UserID  uint   `gorm:"not null; index"`
}
