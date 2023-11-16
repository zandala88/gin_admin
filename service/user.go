package service

import (
	"gin_admin/define"
	"gin_admin/helper"
	"gin_admin/models"
	"github.com/gin-gonic/gin"
)

// UserInfo
// @Tags 登录信息认证-方法
// @Summary 获取用户信息
// @Param AccessToken header string true "AccessToken"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/info [get]
func UserInfo(c *gin.Context) {
	userClaim := c.MustGet("UserClaim").(*define.UserClaim)
	data := new(UserInfoReply)

	err := models.GetUserInfo(userClaim.Identity).Find(data).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "网络异常",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "加载成功",
		"data": data,
	})
}

// UserPasswordChange
// @Tags 登录信息认证-方法
// @Summary 修改密码
// @Param AccessToken header string true "AccessToken"
// @Param password body UserPasswordChangeRequest true "新旧密码"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/password/change [put]
func UserPasswordChange(c *gin.Context) {
	in := new(UserPasswordChangeRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		helper.Error("[BindJSON ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	userClaim := c.MustGet("UserClaim").(*define.UserClaim)

	// 判断旧密码是否正确
	var cnt int64
	err = models.DB.Model(new(models.UserBasic)).
		Where("identity = ? AND password = ?", userClaim.Identity, helper.Md5(in.OldPassword)).Count(&cnt).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "网络异常",
		})
		return
	}
	if cnt == 0 {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "旧密码不正确",
		})
		return
	}

	// 修改密码
	err = models.DB.Model(new(models.UserBasic)).Where("identity = ?", userClaim.Identity).
		Updates(map[string]interface{}{"password": helper.Md5(in.NewPassword)}).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "网络异常",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}

// SetUserList
// @Tags 鉴权接口-方法
// @Summary 管理员列表
// @Param AccessToken header string true "AccessToken"
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /set/user/list [get]
func SetUserList(c *gin.Context) {
	in := &SetUserListRequest{NewQueryRequest()}
	err := c.ShouldBindQuery(in)
	if err != nil {
		helper.Info("[INPUT ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	var (
		cnt  int64
		list = make([]*SetUserListReply, 0)
	)
	err = models.GetUserList(in.Keyword).Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error
	if err != nil {
		helper.Info("[DB ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	for _, v := range list {
		v.CreatedAt = helper.RFC3339ToNormalTime(v.CreatedAt)
		v.UpdatedAt = helper.RFC3339ToNormalTime(v.UpdatedAt)
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "加载成功",
		"data": gin.H{
			"list":  list,
			"count": cnt,
		},
	})
}

// SetUserAdd
// @Tags 鉴权接口-方法
// @Summary 管理员创建
// @Param AccessToken header string true "AccessToken"
// @Param UserAddRequest body SetUserAddRequest true "新增管理员信息"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /set/user/add [post]
func SetUserAdd(c *gin.Context) {
	in := new(SetUserAddRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		helper.Info("[INPUT ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	// 1. 判断用户名是否存在
	var cnt int64
	err = models.DB.Model(new(models.UserBasic)).Where("username = ?", in.Username).Count(&cnt).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	if cnt > 0 {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "用户已存在",
		})
		return
	}

	// 2. 创建用户数据
	err = models.DB.Create(&models.UserBasic{
		Identity:     helper.UUID(),
		RoleIdentity: in.RoleIdentity,
		Username:     in.Username,
		Password:     helper.Md5(in.Password),
		Phone:        in.Phone,
	}).Error
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
		"msg":  "创建成功",
	})
}

// SetUserUpdate
// @Tags 鉴权接口-方法
// @Summary 管理员修改
// @Param AccessToken header string true "AccessToken"
// @Param UserAddRequest body SetUserUpdateRequest true "新增管理员信息"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /set/user/update [put]
func SetUserUpdate(c *gin.Context) {
	in := new(SetUserUpdateRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		helper.Info("[INPUT ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	// 1. 判断用户名是否已存在
	var cnt int64
	err = models.DB.Model(new(models.UserBasic)).Where("identity != ? AND username = ?", in.Identity, in.Username).Count(&cnt).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	if cnt > 0 {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "用户已存在",
		})
		return
	}
	// 2. 修改数据
	err = models.DB.Model(new(models.UserBasic)).Where("identity = ?", in.Identity).Updates(map[string]any{
		"role_identity": in.RoleIdentity,
		"username":      in.Username,
		"phone":         in.Phone,
	}).Error
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
		"msg":  "修改成功",
	})
}

// SetUserDelete
// @Tags 鉴权接口-方法
// @Summary 管理员删除
// @Param AccessToken header string true "AccessToken"
// @Param identity query string true "identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /set/user/delete [delete]
func SetUserDelete(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		})
		return
	}
	// 删除管理员
	err := models.DB.Where("identity = ?", identity).Delete(new(models.UserBasic)).Error
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
		"msg":  "删除成功",
	})
}
