package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID     int    `db:"id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	Price  int    `db:"price"`
}

func main() {
	db, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_lang_study?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接失败", err)
		panic(err)
	}

	// 创建books表, 其中price作为索引
	db.MustExec(`CREATE TABLE IF NOT EXISTS books (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(80) NOT NULL,
		author VARCHAR(30) NOT NULL,
		price INT NOT NULL,
		INDEX idx_price (price))`)

	db.MustExec(`Insert INTO books (title, author, price) VALUES
		('天龙八部', '金庸', 80),
		('仙剑奇侠传', '姚壮宪', 46),
		('三体', '刘慈欣', 90)`)

	// 查询价格大于50的所有书籍
	books := []Book{}
	db.Select(&books, "SELECT * FROM books WHERE price > ?", 50)
	fmt.Println("查询结果：", books)

}
