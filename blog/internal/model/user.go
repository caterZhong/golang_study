package model

// 用户db模型
type User struct {
	ID uint `gorm:"primarykey"`
	// gorm.Model already includes ID, CreatedAt, UpdatedAt, DeletedAt
	UserName string `gorm:"type:varchar(128); not null; index"`
	Password string `gorm:"type:varchar(512); not null;"`
	Email    string `gorm:"type:varchar(128); not null; index"`
}
