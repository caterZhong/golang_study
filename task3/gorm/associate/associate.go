package main

import (
	"fmt"
	"golang_study/task3/gorm/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// fmt.Println("GORM 关联关系示例")

	dsn := "root:root@tcp(localhost:3306)/go_lang_study?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 查询某个用户的所有文章,以王五为例
	fmt.Println("查询某个用户的所有文章,以王五为例")
	var user model.User
	reuslt := db.Preload("Post").Find(&user, "name = ?", "王五")

	if reuslt.Error != nil {
		fmt.Println("查询用户失败", reuslt.Error)
	} else {
		fmt.Printf("用户 %s 的所有文章:\n", user.Name)
		for _, post := range user.Post {
			fmt.Printf("文章ID: %d, 标题: %s, 内容: %s, 评论数: %d\n", post.ID, post.Title, post.Content, post.CommentAmount)
		}
	}

	// 查询评论最多的文章
	fmt.Println("查询评论最多的文章")
	var topPost model.Post
	result2 := db.Order("comment_amount DESC").First(&topPost)
	if result2.Error != nil {
		fmt.Println("查询评论最多的文章失败", result2.Error)
	} else {
		fmt.Printf("评论最多的文章是: 文章ID: %d, 标题: %s, 内容: %s, 评论数: %d\n", topPost.ID, topPost.Title, topPost.Content, topPost.CommentAmount)
	}

}
