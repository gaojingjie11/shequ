package main

import (
	"fmt"
	"os"
	"smartcommunity/internal/config" // 引入新写的 config 包
	"smartcommunity/internal/global"
	"smartcommunity/internal/model"
	"smartcommunity/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// cmd/main.go
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	config.Init(env)

	// 2. 初始化数据库和Redis
	// 直接从 config.Conf 中获取 yaml 里的值
	global.InitDB(config.Conf.DB.DSN)
	global.InitRedis(config.Conf.Redis.Addr, "") // 假设yaml里没配密码，暂时传空

	// 自动迁移
	global.DB.AutoMigrate(&model.SysUser{})

	// 3. 启动 Gin
	r := gin.Default()
	router.InitRouter(r)

	fmt.Println("服务启动在 :8080")
	r.Run(":8080")
}
