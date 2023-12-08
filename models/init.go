package models

import (
	"github.com/go-ini/ini"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

var RDB *redis.Client

func NewGormDB() {
	file, err := ini.Load("conf/config.ini")
	if err != nil {
		panic(err)
	}
	host := file.Section("mysql").Key("MySQLHost").String()
	port := file.Section("mysql").Key("MySQLPort").String()
	user := file.Section("mysql").Key("MySQLUser").String()
	password := file.Section("mysql").Key("MySQLPassword").String()
	dbname := file.Section("mysql").Key("MySQLDBName").String()

	addr := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	log.Println(addr)
	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&UserBasic{}, &RoleBasic{}, &RoleMenu{}, &RoleFunction{}, &MenuBasic{}, &FunctionBasic{})
	if err != nil {
		panic(err)
	}
	DB = db
}

func NewRedisDB() {
	file, err := ini.Load("conf/config.ini")
	if err != nil {
		panic(err)
	}
	RDB = redis.NewClient(&redis.Options{
		Addr:     file.Section("redis").Key("RedisHost").String() + ":" + file.Section("redis").Key("RedisPort").String(),
		Username: file.Section("redis").Key("RedisUsername").String(),
		Password: file.Section("redis").Key("RedisPassword").String(),
	})
}
