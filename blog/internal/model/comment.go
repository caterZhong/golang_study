package model

import (
	"gorm.io/gorm"
)

// 评论db模型
type Comment struct {
	gorm.Model
	Content string `gorm:"type:text; not null"`
	PostID  uint   `gorm:"not null; index"`
	UserID  uint   `gorm:"not null; index"`
}
