package test

import (
	"context"
	"fmt"
	"gin_admin/models"
	"github.com/go-ini/ini"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestMySQL(t *testing.T) {
	file, err := ini.Load("../conf/config.ini")
	if err != nil {
		panic(err)
	}
	host := file.Section("mysql").Key("MySQLHost").String()
	port := file.Section("mysql").Key("MySQLPort").String()
	user := file.Section("mysql").Key("MySQLUser").String()
	password := file.Section("mysql").Key("MySQLPassword").String()
	dbname := file.Section("mysql").Key("MySQLDBName").String()

	addr := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.RoleBasic, 0)
	err = db.Find(&data).Error
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range data {
		fmt.Printf("%+v\n", v)
	}
}

func TestRedis(t *testing.T) {
	var ctx = context.Background()
	file, err := ini.Load("../conf/config.ini")
	if err != nil {
		panic(err)
	}
	var rdb = redis.NewClient(&redis.Options{
		Addr:     file.Section("redis").Key("RedisHost").String() + ":" + file.Section("redis").Key("RedisPort").String(),
		Username: file.Section("redis").Key("RedisUsername").String(),
		Password: file.Section("redis").Key("RedisPassword").String(),
	})
	rdb.Set(ctx, "test", "test", time.Second*10)
	v, err := rdb.Get(ctx, "test").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(v)

}
