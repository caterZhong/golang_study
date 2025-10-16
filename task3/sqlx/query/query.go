package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

func main() {
	db, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_lang_study?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接失败", err)
		panic(err)
	}

	// 创建员工表, 其中salary作为索引
	db.MustExec(`CREATE TABLE IF NOT EXISTS employees (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(30) NOT NULL,
		department VARCHAR(30) NOT NULL,
		salary INT NOT NULL,
		INDEX idx_salary (salary))`)

	db.MustExec(`Insert INTO employees (name, department, salary) VALUES
		('张三', '技术部', 8000),
		('李四', '市场部', 6000),
		('王五', '技术部', 9000)`)

	// 查询技术部的所有员工
	employees := []Employee{}
	db.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部")
	fmt.Println("查询结果：", employees)

	// 查询做高工资的员工信息
	var topEployee Employee
	db.Get(&topEployee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	fmt.Println("工资最高的员工信息：", topEployee)

}
