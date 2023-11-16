package router

import (
	_ "gin_admin/docs"
	"gin_admin/middleware"
	"gin_admin/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func App() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	// Swagger 配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 用户名、密码登录
	r.POST("/login/password", service.LoginPassword)

	// 登录信息认证
	loginAuth := r.Group("/").Use(middleware.LoginAuthCheck())
	{
		// 获取菜单列表
		loginAuth.GET("/menus", service.Menus)
		// 获取功能列表
		loginAuth.GET("/functions", service.Functions)
		// 获取用户信息
		loginAuth.GET("/user/info", service.UserInfo)
		// 修改密码
		loginAuth.PUT("/user/password/change", service.UserPasswordChange)
	}

	// 进行鉴权的接口
	auth := loginAuth.Use(middleware.FuncAuthCheck())
	{
		// ---------------- BEGIN - 设置 ----------------
		// 角色功能管理
		// 角色列表
		auth.GET("/set/role/list", service.SetRoleList)
		// 角色详情
		auth.GET("/set/role/detail", service.SetRoleDetail)
		// 修改角色的管理员身份
		auth.PUT("/set/role/update/admin", service.SetRoleUpdateAdmin)
		// 新增角色
		auth.POST("/set/role/create", service.SetRoleCreate)
		// 删除角色
		auth.DELETE("/set/role/delete", service.SetRoleDelete)
		// 修改角色
		auth.PUT("/set/role/update", service.SetRoleUpdate)

		// 菜单功能管理
		// 菜单列表
		auth.GET("/set/menu/list", service.SetMenuList)
		// 功能列表
		auth.GET("/set/func/list", service.SetFuncList)

		// 管理员管理
		// 管理员列表
		auth.GET("/set/user/list", service.SetUserList)
		// 新增管理员
		auth.POST("/set/user/add", service.SetUserAdd)
		// 修改管理员
		auth.PUT("/set/user/update", service.SetUserUpdate)
		// 删除管理员
		auth.DELETE("/set/user/delete", service.SetUserDelete)
		// ---------------- END - 设置 ----------------

		// ---------------- BEGIN - dev ----------------
		// 新增菜单
		auth.POST("/dev/menu/add", service.DevMenuAdd)
		// ---------------- END - dev ----------------
	}
	return r
}
