package main

import (
	"fmt"
	"golang_study/blog/config"
	"golang_study/blog/internal/api"
	"golang_study/blog/internal/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.ReadConfig("../config/config.yaml")
	if err != nil {
		log.Fatalf("config load failed: %v", err)
	}

	// 初始化MySQL连接池
	if err := config.InitDB(&cfg.DB); err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}

	// 初始化jwt
	if err := utils.NewJWTService(&cfg.JWT); err != nil {
		fmt.Println(err)
		log.Fatalf("jwt初始化失败: %v", err)
	}

	// 初始化Redis连接池
	// if err := redis.Init(cfg); err != nil {
	//     log.Fatalf("Redis初始化失败: %v", err)
	// }

	// 用户模块的di初始化

	// post模块的di初始化

	// comment模块的di初始化

	fmt.Println(utils.GenHashedPass("test@163.com"))

	fmt.Println("main 函数开始了")
	fmt.Println("开始注册路由")
	r := gin.Default()
	api.RegisterRoutes(r)

	fmt.Println("开始启动服务")

	r.Run()
}
