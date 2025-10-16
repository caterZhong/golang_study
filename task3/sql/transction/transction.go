package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	ID        uint
	Balance   uint      `gorm:"not null, default:0"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type Transaction struct {
	ID            uint
	FromAccountId uint `gorm:"not null"`
	ToAccountId   uint `gorm:"not null"`
	Amount        uint `gorm:"not null"`
}

func main() {

	dsn := "root:root@tcp(localhost:3306)/go_lang_study?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Account{}, &Transaction{})

	// 创建两个账户
	account1 := Account{Balance: 1000}
	account2 := Account{Balance: 500}
	db.Create(&account1)
	db.Create(&account2)

	// 执行转账操作
	err = db.Transaction(func(tx *gorm.DB) error {
		// 检查账户余额
		var amount uint = 100
		var fromAccount Account = account1
		if fromAccount.Balance < amount {
			return fmt.Errorf("余额不足")
		}
		// 扣款

		// 生成转账记录
		transaction := Transaction{
			FromAccountId: account1.ID,
			ToAccountId:   account2.ID,
			Amount:        100,
		}
		if err := tx.Debug().Create(&transaction).Error; err != nil {
			return err
		}

		// 扣款
		if err := tx.Debug().Model(&Account{}).Where("id = ? and balance > ?", account1.ID, amount).Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return err
		}

		// 加款
		if err := tx.Debug().Model(&Account{}).Where("id = ?", account2.ID).Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println("转账失败:", err)
	}

	fmt.Println("转账成功")

}
