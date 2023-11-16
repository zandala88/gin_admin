package main

import (
	"gin_admin/define"
	"gin_admin/models"
	"gin_admin/router"
)

func main() {
	// 初始化 gorm.DB
	models.NewGormDB()
	// 初始化 redis.Client
	models.NewRedisDB()
	r := router.App()
	r.Run(define.Port)
}
