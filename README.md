# gin_admin

## go配置

```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```



## 依赖安装

```shell
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite

go get -u github.com/gin-gonic/gin
go mod init github.com/my/repo

go get github.com/redis/go-redis/v9

go get -u github.com/golang-jwt/jwt/v5

go get github.com/satori/go.uuid

go get github.com/go-ini/ini
go get gopkg.in/ini.v1
```



## 配置
1. 创建数据库，将数据库名置为 `up-admin` ,并执行 `./sql/*.sql` 文件
2. 修改 `conf/config.ini` 中的 
	- redis
    - mysql
3. 删除 `docs` 文件夹后运行

	```shell 
	go get -u github.com/swaggo/swag/cmd/swag
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/swaggo/gin-swagger
	go get -u github.com/swaggo/files
	
	swag init
	```
	
	

## Swagger 接口文档

### 访问地址 http://127.0.0.1:9090/swagger/index.html



## 解析

登录接口：生成token，放置Header中用于后续请求

管理员账号 `get` 密码 `123456`

