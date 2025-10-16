package model

import (
	"gorm.io/gorm"
)

type User struct {
	ID         uint
	Name       string `gorm:"type:varchar(128); not null; index"`
	PostAmount uint
	Post       []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	ID      uint
	Title   string `gorm:"type:varchar(256); not null; index"`
	Content string `gorm:"type:text; not null"`
	// 外键
	UserID        uint   `gorm:"not null;"`
	CommentAmount uint   `gorm:"not null; default:0; index"`
	CommentStatus string `gorm:"type:varchar(16); not null; default:'无评论'"`
}

type Comment struct {
	ID      uint
	Content string `gorm:"type:text; not null"`
	// 外键
	PostID uint `gorm:"not null;"`
	Post   Post `gorm:"foreignKey:PostID"`
	UserID uint `gorm:"not null; index"`
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 每次创建文章后，自动更新用户的文章数量
	if err := tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_amount", gorm.Expr("post_amount + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	// 每次新增评论后，检查该评论所属文章的评论数量是否为0 ， 如果为0， 则更新文章评论状态为“无评论”
	tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_amount", gorm.Expr("comment_amount + ?", 1)).Error; err != nil {
			return err
		}

		var post Post
		result := tx.First(&post, c.PostID)

		if result.Error != nil {
			return result.Error
		}

		var commentStutus string = "无评论"
		if post.CommentAmount > 0 {
			commentStutus = "有评论"
		}

		if err := tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", commentStutus).Error; err != nil {
			return err
		}

		return nil
	})

	return nil
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// 每次删除评论后，检查该评论所属文章的评论数量是否为0 ， 如果为0， 则更新文章评论状态为“无评论”
	tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Post{}).Where("id = ? and comment_status > 0", c.PostID).Update("comment_amount", gorm.Expr("comment_amount - ?", 1)).Error; err != nil {
			return err
		}

		var post Post
		result := tx.First(&post, c.PostID)

		if result.Error != nil {
			return result.Error
		}

		var commentStutus string = "无评论"
		if post.CommentAmount > 0 {
			commentStutus = "有评论"
		}

		if err := tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", commentStutus).Error; err != nil {
			return err
		}

		return nil
	})

	return nil
}

// func addPost(db *gorm.DB, userId uint, posts *[]Post) int {
// 	for index, post := range *posts {
// 		(*posts)[index].UserID = userId
// 		if post.Title == "" {
// 			panic("文章标题不能为空")
// 		}
// 		if post.Content == "" {
// 			panic("文章内容不能为空")
// 		}
// 	}

// 	err := db.Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Debug().Create(&posts).Error; err != nil {
// 			return err
// 		}

// 		if err := tx.Debug().Model(&User{}).Where("id = ?", userId).Update("post_amount", gorm.Expr("post_amount + ?", len(*posts))).Error; err != nil {
// 			return err
// 		}

// 		return nil
// 	})

// 	if err != nil {
// 		fmt.Println("添加文章失败", err)
// 		return 0
// 	}
// 	return len(*posts)

// }

// func addComment(db *gorm.DB, userId uint, postId uint, comments *[]Comment) int {
// 	for index, comment := range *comments {
// 		(*comments)[index].PostID = postId
// 		(*comments)[index].UserID = userId
// 		if comment.Content == "" {
// 			panic("评论内容不能为空")
// 		}
// 	}
// 	err := db.Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Debug().Create(&comments).Error; err != nil {
// 			return err
// 		}

// 		if err := tx.Debug().Model(&Post{}).Where("id = ?", postId).Update("comment_amount", gorm.Expr("comment_amount + ?", len(*comments))).Error; err != nil {
// 			return err
// 		}

// 		return nil
// 	})

// 	if err != nil {
// 		fmt.Println("添加评论失败", err)
// 		return 0
// 	}
// 	return len(*comments)

// }

// func initializeData(db *gorm.DB) {
// 	db.Debug().AutoMigrate(&User{}, &Post{}, &Comment{})
// 	// 初始化数据
// 	users := []User{
// 		{Name: "张三"},
// 		{Name: "李四"},
// 		{Name: "王五"},
// 	}
// 	result := db.Debug().Create(&users)
// 	if result.Error != nil {
// 		panic("插入用户失败")
// 	}

// 	// 批量张三的文章
// 	posts := []Post{
// 		{Title: "张三的第一篇文章", Content: "这是张三的第一篇文章", UserID: users[0].ID},
// 		{Title: "张三的第二篇文章", Content: "这是张三的第二篇文章", UserID: users[0].ID},
// 	}

// 	addedCount := addPost(db, users[0].ID, &posts)
// 	fmt.Printf("成功添加了 %s 的 %d 篇文章\n", users[0].Name, addedCount)

// 	// 给张三第一篇文章添加1条评论
// 	comments1 := []Comment{
// 		{Content: "这是张三的第一篇文章的第一条评论"},
// 	}
// 	fmt.Println("users:", users)
// 	fmt.Println("posts:", posts)
// 	addedCommentCount1 := addComment(db, users[0].ID, posts[0].ID, &comments1)
// 	fmt.Printf("成功添加了 %s 的 %d 条评论\n", users[0].Name, addedCommentCount1)

// 	// 给张三第二篇文章添加2条评论
// 	comments2 := []Comment{
// 		{Content: "这是张三的第二篇文章的第一条评论"},
// 		{Content: "这是张三的第二篇文章的第二条评论"},
// 	}
// 	addedCommentCount2 := addComment(db, users[0].ID, posts[1].ID, &comments2)
// 	fmt.Printf("成功添加了 %s 的 %d 条评论\n", users[0].Name, addedCommentCount2)

// 	// 添加李四的1篇文章
// 	posts2 := []Post{
// 		{Title: "李四的第一篇文章", Content: "这是李四的第一篇文章", UserID: users[1].ID},
// 	}
// 	addedCount2 := addPost(db, users[1].ID, &posts2)
// 	fmt.Printf("成功添加了 %s 的 %d 篇文章\n", users[1].Name, addedCount2)

// 	// 给李四的文章添加3条评论
// 	comments3 := []Comment{
// 		{Content: "这是李四的第一篇文章的第一条评论"},
// 		{Content: "这是李四的第一篇文章的第二条评论"},
// 		{Content: "这是李四的第一篇文章的第三条评论"},
// 	}
// 	addedCommentCount3 := addComment(db, users[1].ID, posts2[0].ID, &comments3)
// 	fmt.Printf("成功添加了 %s 的 %d 条评论\n", users[1].Name, addedCommentCount3)

// 	// 添加王五的3篇文章
// 	post3 := []Post{
// 		{Title: "王五的第一篇文章", Content: "这是王五的第一篇文章", UserID: users[2].ID},
// 		{Title: "王五的第二篇文章", Content: "这是王五的第二篇文章", UserID: users[2].ID},
// 		{Title: "王五的第三篇文章", Content: "这是王五的第三篇文章", UserID: users[2].ID},
// 	}
// 	addedCount3 := addPost(db, users[2].ID, &post3)
// 	fmt.Printf("成功添加了 %s 的 %d 篇文章\n", users[2].Name, addedCount3)

// 	// 给王五的3篇文章分别添加1、2、5条评论
// 	comments4 := []Comment{
// 		{Content: "这是王五的第一篇文章的第一条评论"},
// 	}
// 	addedCommentCount4 := addComment(db, users[2].ID, post3[0].ID, &comments4)
// 	fmt.Printf("成功添加了 %s 的 %d 条评论\n", users[2].Name, addedCommentCount4)

// 	comments5 := []Comment{
// 		{Content: "这是王五的第二篇文章的第一条评论"},
// 		{Content: "这是王五的第二篇文章的第二条评论"},
// 	}
// 	addedCommentCount5 := addComment(db, users[2].ID, post3[1].ID, &comments5)
// 	fmt.Printf("成功添加了 %s 的 %d 条评论\n", users[2].Name, addedCommentCount5)

// 	comments6 := []Comment{
// 		{Content: "这是王五的第三篇文章的第一条评论"},
// 		{Content: "这是王五的第三篇文章的第二条评论"},
// 		{Content: "这是王五的第三篇文章的第三条评论"},
// 		{Content: "这是王五的第三篇文章的第四条评论"},
// 		{Content: "这是王五的第三篇文章的第五条评论"},
// 	}
// 	addedCommentCount6 := addComment(db, users[2].ID, post3[2].ID, &comments6)
// 	fmt.Printf("成功添加了 %s 的 %d 条评论\n", users[2].Name, addedCommentCount6)

// 	fmt.Println("数据初始化完成")

// }

// func main() {

// 	dsn := "root:root@tcp(localhost:3306)/go_lang_study?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	initializeData(db)

// 	// 查询某个用户的所有文章,以王五为例
// 	fmt.Println("查询某个用户的所有文章,以王五为例")
// 	var user User
// 	reuslt := db.Preload("Post").Find(&user, "name = ?", "王五")

// 	if reuslt.Error != nil {
// 		fmt.Println("查询用户失败", reuslt.Error)
// 	} else {
// 		fmt.Printf("用户 %s 的所有文章:\n", user.Name)
// 		for _, post := range user.Post {
// 			fmt.Printf("文章ID: %d, 标题: %s, 内容: %s, 评论数: %d\n", post.ID, post.Title, post.Content, post.CommentAmount)
// 		}
// 	}

// 	// 查询评论最多的文章
// 	fmt.Println("查询评论最多的文章")
// 	var topPost Post
// 	result2 := db.Order("comment_amount DESC").First(&topPost)
// 	if result2.Error != nil {
// 		fmt.Println("查询评论最多的文章失败", result2.Error)
// 	} else {
// 		fmt.Printf("评论最多的文章是: 文章ID: %d, 标题: %s, 内容: %s, 评论数: %d\n", topPost.ID, topPost.Title, topPost.Content, topPost.CommentAmount)
// 	}

// }
