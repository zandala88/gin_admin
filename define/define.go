package define

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// MySQL配置
// username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
var UpAdminDSN = "root:root@tcp(127.0.0.1:3306)/up-admin?charset=utf8mb4&parseTime=True&loc=Local"

var (
	Port           = ":9090"
	FrameName      = "GinAdmin"
	DateTimeLayout = "2006-01-02 15:04:05"
	// DefaultSize 默认每页查询20条数据
	DefaultSize = 20
	// JwtKey 密钥
	JwtKey = "gin-admin"
	// TokenExpire token 有效期，7天
	TokenExpire = time.Now().Add(time.Second * 3600 * 24 * 7).Unix()
)

// Redis 配置
var (
	RedisAddr     = "localhost:6379"
	RedisUsername = ""
	RedisPassword = ""
	// RedisRoleAdminPrefix 判断角色是否是超管的前缀
	RedisRoleAdminPrefix = "ADMIN-"
	// RedisMenuPrefix 菜单
	RedisMenuPrefix = "MENU"
	// RedisFuncPrefix 功能的前缀
	RedisFuncPrefix = "FUNC-"
)

type UserClaim struct {
	Id           uint
	Identity     string
	Name         string
	RoleIdentity string // 角色唯一标识
	IsAdmin      bool   // 是否是超管
	jwt.StandardClaims
}
