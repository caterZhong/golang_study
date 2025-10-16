package main

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	ID        uint
	Name      string `gorm:"index;type:varchar(128);not null"`
	Age       int8   `gorm:"not null"`
	Grade     string `gorm:"type:varchar(16); not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted int8 `gorm:"default:0; not null"`
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&Student{})
}

func Insert(db *gorm.DB) {
	// student := Student{
	// 	Name:  "张三",
	// 	Age:   18,
	// 	Grade: "一年级",
	// }
	// db.Create(&student)
	students := []Student{{
		Name:  "李四",
		Age:   20,
		Grade: "二年级"}, {
		Name:  "王五",
		Age:   22,
		Grade: "三年级"},
	}
	db.Create(&students)
}

func Query(db *gorm.DB) {
	student := Student{}
	result := db.First(&student, 1)
	result = db.First(&student, "name = ?", "李四")
	fmt.Println("查询结果：", result.RowsAffected)
	hasFound := errors.Is(result.Error, gorm.ErrRecordNotFound)
	fmt.Println("是否没有找到数据", hasFound)
	fmt.Println("查询结果内容：", student)

	student2 := Student{ID: 3, Age: 30}
	db.Debug().Model(Student{Name: "李四"}).First(&student2, "name = ?", "20")
	fmt.Println("Model First 查询new结果内容：", student2)

	var students []Student
	db.Debug().Where("name = ?", "李四").Find(&student2, "name = ?", "王五")
	fmt.Println("Where First 查询结果内容：", student2)

	db.Debug().Where("name <> ?", "李四").Find(&students)
	fmt.Println("Where Find 查询结果内容：", students)

	db.Debug().Where("name LIKE ?", "%李%").Find(&students)
	fmt.Println("Where Like 查询结果内容：", students)

	// IN 查询
	db.Debug().Where("id IN ?", []int{4, 5, 3}).Find(&students)
	fmt.Println("Where In 查询结果内容：", students)

	// AND
	db.Debug().Where("name = ? AND age = ?", "李四", 20).Find(&students)
	fmt.Println("Where AND 查询结果内容：", students)

	// Time
	db.Debug().Where("created_at > ?", time.Now()).Find(&students)
	fmt.Println("Where Time 查询结果内容：", students)

	// Between
	db.Debug().Where("created_at BETWEEN ? AND ?", "2024-06-01", "2024-06-30").Find(&students)
	fmt.Println("Where Between 查询结果内容：", students)

	// Is Null
	db.Debug().Where("updated_at IS NULL").Find(&students)
	fmt.Println("Where Is Null 查询结果内容：", students)

	// struct
	db.Debug().Where(&Student{Age: 21}).Find(&students)
	fmt.Println("Where Struct 查询结果内容：", students)

	// map
	db.Debug().Where(map[string]interface{}{"age": 21}).Find(&students)
	fmt.Println("Where Map 查询结果内容：", students)

	// 结构体
	db.Debug().Where(Student{Name: "jinzhu", Age: 0}).Find(&students)
	fmt.Println("结构体查询结果内容：", students)

	db.Debug().Where(&Student{Name: "jinzhu", Age: 1}, "name", "Age").Find(&students)
	fmt.Println("结构体指定字段查询结果内容：", students)

	db.Debug().Where(&Student{Name: "jinzhu"}, "Age").Find(&students)
	fmt.Println("结构体指定字段2查询结果内容：", students)

	db.Debug().Not("name", "张三").Find(&students)
	fmt.Println("Not 查询结果内容：", students)

	// or  + 结构体
	db.Debug().Where("name = 'jinzhu'").Or(&Student{Name: "jinzhu 2", Age: 18}).Find(&students)
	fmt.Println("Or + 结构体 查询结果内容：", students)

	// or  + 结构体
	db.Debug().Find(&students, "name = 'jin' or age = 3")
	fmt.Println("Or 内联 查询结果内容：", students)

	// db.Debug().Table("students").Select("COALESCE(age,?)", 42).Rows()
	// fmt.Println("COALESCE 查询结果内容：", age)

	resultBatch := db.Find(&students, []int{1, 2, 3})
	fmt.Println("是否没有找到数据", errors.Is(resultBatch.Error, gorm.ErrRecordNotFound))
	fmt.Println("查询结果数据：", resultBatch.RowsAffected)
	fmt.Println("查询结果内容：", len(students))
}

func Update(db *gorm.DB) {
	db.AutoMigrate(&Student{})
}

func Delete(db *gorm.DB) {
	db.AutoMigrate(&Student{})
}

func (s *Student) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("数据创建前的钩子函数")
	fmt.Println(s)
	return
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_lang_study?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	// db.Debug().AutoMigrate(&Student{})
	// Insert(db)
	Query(db)
	// db.Debug().AutoMigrate(&Student{})
	// db.Debug().AutoMigrate(&Student{})
	// db.Debug().AutoMigrate(&Student{})

}
