# gin_admin

## 依赖安装
```shell
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u github.com/gin-gonic/gin
go mod init github.com/my/repo
go get github.com/redis/go-redis/v9
go get -u github.com/golang-jwt/jwt/v5
go get github.com/satori/go.uuid
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
swag init
```



## 配置
1. 创建数据库，将数据库名置为 `up-admin` ,并执行 `.sql` 文件
2. 修改 `define.go` 中的 
- `RedisAddr` 
- `RedisUsername`
- `RedisPassword`
- `UpAdminDSN (dbname = up-admin)`




## Swagger 接口文档
### 访问地址 http://127.0.0.1:9090/swagger/index.html




## 功能
- [x] 登录
