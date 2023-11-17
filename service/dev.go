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

// DevMenuUpdate
// @Tags 鉴权接口-方法
// @Summary 修改菜单
// @Param AccessToken header string true "AccessToken"
// @Param MenuUpdateRequest body DevMenuUpdateRequest true "修改菜单信息"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /dev/menu/update [put]
func DevMenuUpdate(c *gin.Context) {
	in := new(DevMenuUpdateRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		helper.Error("[BindJSON ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	if in.Identity == "" {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
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
	// 2. 更新数据
	err = models.DB.Model(new(models.MenuBasic)).Where("identity = ?", in.Identity).Updates(map[string]interface{}{
		"parent_id": parentId,
		"name":      in.Name,
		"web_icon":  in.WebIcon,
		"sort":      in.Sort,
		"path":      in.Path,
		"level":     in.Level,
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
		"msg":  "修改成功",
	})
}

// DevMenuDelete
// @Tags 鉴权接口-方法
// @Summary 删除菜单
// @Param AccessToken header string true "AccessToken"
// @Param identity query string true "identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /dev/menu/delete [delete]
func DevMenuDelete(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		})
		return
	}
	// 1. 删除数据库中的数据
	err := models.DB.Where("identity = ?", identity).Delete(new(models.MenuBasic)).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	// 2. 清空菜单缓存数据
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
		"msg":  "删除成功",
	})
}

// DevFuncAdd
// @Tags 鉴权接口-方法
// @Summary 新增功能
// @Param AccessToken header string true "AccessToken"
// @Param FuncAddRequest body DevFuncAddRequest true "新增功能信息"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /dev/func/add [post]
func DevFuncAdd(c *gin.Context) {
	in := new(DevFuncAddRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		helper.Error("[BindJSON ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	if in.MenuIdentity == "" {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		})
		return
	}
	// 1. 获取菜单ID
	var menuId uint
	err = models.DB.Model(new(models.MenuBasic)).Select("id").
		Where("identity = ?", in.MenuIdentity).Scan(&menuId).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	// 2. 保存数据
	err = models.DB.Create(&models.FunctionBasic{
		Identity: helper.UUID(),
		MenuId:   menuId,
		Name:     in.Name,
		Uri:      in.Uri,
		Sort:     in.Sort,
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
		"msg":  "新增成功",
	})
}

// DevFuncUpdate
// @Tags 鉴权接口-方法
// @Summary 修改功能
// @Param AccessToken header string true "AccessToken"
// @Param FuncUpdateRequest body DevFuncUpdateRequest true "修改功能信息"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /dev/func/update [put]
func DevFuncUpdate(c *gin.Context) {
	in := new(DevFuncUpdateRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		helper.Error("[BindJSON ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	if in.Identity == "" || in.MenuIdentity == "" {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		})
		return
	}
	// 1. 获取菜单ID
	var menuId uint
	err = models.DB.Model(new(models.MenuBasic)).Select("id").
		Where("identity = ?", in.MenuIdentity).Scan(&menuId).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	// 2. 更新数据
	err = models.DB.Model(new(models.FunctionBasic)).Where("identity = ?", in.Identity).Updates(map[string]interface{}{
		"menu_id": menuId,
		"name":    in.Name,
		"uri":     in.Uri,
		"sort":    in.Sort,
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

// DevFuncDelete
// @Tags 鉴权接口-方法
// @Summary 删除功能
// @Param AccessToken header string true "AccessToken"
// @Param identity query string true "identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /dev/func/delete [delete]
func DevFuncDelete(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		})
		return
	}
	err := models.DB.Where("identity = ?", identity).Delete(new(models.FunctionBasic)).Error
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
