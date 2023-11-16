package helper

import (
	"crypto/md5"
	"errors"
	"fmt"
	"gin_admin/define"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"strings"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// GenerateToken
// 生成token
func GenerateToken(id uint, identity, name, roleIdentity string, expireAt int64) (string, error) {
	uc := define.UserClaim{
		Id:           id,
		Identity:     identity,
		Name:         name,
		RoleIdentity: roleIdentity,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyzeToken
// Token 解析
func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
}

// RFC3339ToNormalTime
// RFC3339 日期格式标准化
func RFC3339ToNormalTime(rfc3339 string) string {
	if len(rfc3339) < 19 || rfc3339 == "" || !strings.Contains(rfc3339, "T") {
		return rfc3339
	}
	return strings.Split(rfc3339, "T")[0] + " " + strings.Split(rfc3339, "T")[1][:8]
}

func UUID() string {
	return uuid.NewV4().String()
}
