package config

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func InitDB(dbconfig *DBConfig) error {

	dsn := dbconfig.User + ":" + dbconfig.Password +
		"@tcp(" + dbconfig.Host + ":" + dbconfig.Port +
		")/" + dbconfig.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"

	var initErr error

	once.Do(func() {
		sqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			initErr = err
			return
		}

		// 获取通用数据库对象 sql.DB
		rawDB, err := sqlDB.DB()
		if err != nil {
			initErr = err
			return
		}

		// 设置连接池参数
		rawDB.SetMaxOpenConns(dbconfig.MaxConns)
		rawDB.SetMaxIdleConns(dbconfig.MaxIdleConns)
		rawDB.SetConnMaxLifetime(dbconfig.ConnMaxLifetime)
		DB = sqlDB
	})
	return initErr
}
