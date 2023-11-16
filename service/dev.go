package service

import (
	"context"
	"gin_admin/define"
	"gin_admin/helper"
	"gin_admin/models"
	"github.com/gin-gonic/gin"
)

// DevMenuAdd
// @Tags 鉴权接口-方法
// @Summary 新增菜单
// @Param AccessToken header string true "AccessToken"
// @Param MenuAddRequest body DevMenuAddRequest true "新增菜单信息"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /dev/menu/add [post]
func DevMenuAdd(c *gin.Context) {
	in := new(DevMenuAddRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		helper.Error("[BindJSON ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	// 1. 获取父级ID
	var parentId uint
	if in.ParentIdentity != "" {
		err = models.DB.Model(new(models.MenuBasic)).Select("id").
			Where("identity = ?", in.ParentIdentity).Scan(&parentId).Error
		if err != nil {
			helper.Error("[DB ERROR] : %v", err)
			c.JSON(200, gin.H{
				"code": -1,
				"msg":  "数据库异常",
			})
			return
		}
	}
	// 2. 保存数据
	err = models.DB.Create(&models.MenuBasic{
		Identity: helper.UUID(),
		ParentId: parentId,
		Name:     in.Name,
		WebIcon:  in.WebIcon,
		Sort:     in.Sort,
		Path:     in.Path,
		Level:    in.Level,
	}).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	// 3. 清空菜单缓存数据
	err = models.RDB.Del(context.Background(), define.RedisMenuPrefix).Err()
	if err != nil {
		helper.Error("[RDB ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "缓存异常",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "新增成功",
	})
}
