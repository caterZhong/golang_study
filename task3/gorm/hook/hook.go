package main

import (
	"fmt"
	"golang_study/task3/gorm/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(localhost:3306)/go_lang_study?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// posts := []model.Post{
	// 	{Title: "测试张三文章1", Content: "这是测试张三文章1的内容", UserID: 1},
	// 	{Title: "测试张三文章2", Content: "这是测试张三文章2的内容", UserID: 1},
	// }

	// result := db.Create(&posts)
	// if result.Error != nil {
	// 	fmt.Println("添加文章失败", result.Error)
	// } else {
	// 	fmt.Printf("成功添加了 %d 篇文章\n", result.RowsAffected)
	// }

	// // 查询张三的信息， 观察钩子是否生效， 用户发布的文章数量是否生效
	// var user model.User
	// result2 := db.First(&user, 1)
	// if result2.Error != nil {
	// 	fmt.Println("查询用户失败", result2.Error)
	// } else {
	// 	fmt.Printf("用户 %s 共有 %d 篇文章\n", user.Name, user.PostAmount)
	// }

	comments := []model.Comment{
		{Content: "这是测试张三文章1的第一条评论", PostID: 7},
		{Content: "这是测试张三文章1的第二条评论", PostID: 7},
		{Content: "这是测试张三文章1的第三条评论", PostID: 7},
	}

	result3 := db.Create(&comments)
	if result3.Error != nil {
		fmt.Println("添加评论失败", result3.Error)
	} else {
		fmt.Printf("成功添加了 %d 条评论\n", result3.RowsAffected)
	}

}
