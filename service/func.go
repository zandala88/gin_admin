package service

import (
	"gin_admin/define"
	"gin_admin/helper"
	"gin_admin/models"
	"github.com/gin-gonic/gin"
)

// Functions
// @Tags 登录信息认证-方法
// @Summary 获取功能列表
// @Param AccessToken header string true "AccessToken"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /functions [get]
func Functions(c *gin.Context) {
	SetFuncList(c)
}

// SetFuncList
// @Tags 鉴权接口-方法
// @Summary 获取功能列表
// @Param AccessToken header string true "AccessToken"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /set/func/list [get]
func SetFuncList(c *gin.Context) {
	userClaim := c.MustGet("UserClaim").(*define.UserClaim)
	data := make([]*SetFuncListReply, 0)
	err := models.GetRoleFunctions(userClaim.RoleIdentity, userClaim.IsAdmin).Find(&data).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "加载成功",
		"data": data,
	})
}
