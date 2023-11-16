package service

import (
	"errors"
	"gin_admin/define"
	"gin_admin/helper"
	"gin_admin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LoginPassword
// @Tags 公共方法
// @Summary 用户登录
// @Param user body LoginPasswordRequest true "登录信息"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /login/password [post]
func LoginPassword(c *gin.Context) {
	in := new(LoginPasswordRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		helper.Error("[BindJSON ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	// 根据账号、密码查询用户
	ub, err := models.GetUserBasicByUsernamePassword(in.Username, helper.Md5(in.Password))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(200, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			})
			return
		}
	}

	// 生成 token
	token, err := helper.GenerateToken(ub.ID, ub.Identity, ub.Username, ub.RoleIdentity, define.TokenExpire)
	if err != nil {
		helper.Error("[BindJSON ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	// 生成 token
	refreshToken, err := helper.GenerateToken(ub.ID, ub.Identity, ub.Username, ub.RoleIdentity, define.RefreshTokenExpire)
	if err != nil {
		helper.Error("[BindJSON ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	data := &LoginPasswordReply{
		Token:        token,
		RefreshToken: refreshToken,
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": data,
	})
}
