package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	ID        uint
	Name      string    `gorm:"index;type:varchar(128);not null"`
	Age       int8      `gorm:"not null"`
	Grade     string    `gorm:"type:varchar(16); not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func Insert(db *gorm.DB, student *Student) {
	db.Create(student)
}

func Query(db *gorm.DB, student *Student) {
	db.Create(student)
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_lang_study?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// db.AutoMigrate(&Student{})

	// result := db.Create(&Student{
	// 	Name:  "张三fen",
	// 	Age:   12,
	// 	Grade: "三年级"})

	// fmt.Println("插入结果：", result.RowsAffected)

	// db.Create(&Student{
	// 	Name:  "张三",
	// 	Age:   20,
	// 	Grade: "三年级"})

	var students []Student
	// student3 := Student{Name: "张三", Age: 20, Grade: "三年级", ID: 4}
	db.Model(&Student{}).Where("age > ?", 18).Find(&students)
	fmt.Println("查询结果：", students)

	db.Debug().Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")

	db.Debug().Where("age < ?", 15).Delete(&Student{})

}
