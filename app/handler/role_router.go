package controller

import (
	"github.com/cilidm/base-system-v2/app/middleware"
	"github.com/cilidm/base-system-v2/app/router"
)

func init() {
	r := router.New("system", "/system", middleware.AuthMiddleware)
	r.GET("admin/list", AdminList) // 管理员列表页
	r.GET("admin/list_json", AdminListJson)
	r.GET("admin/add", AdminAdd)
	r.POST("admin/add", AdminAddHandler)
	r.GET("admin/edit", AdminEdit)
	r.POST("admin/edit", AdminEditHandler)
	r.POST("admin/delete", AdminDelete)

	r.GET("role/list", RoleList) // 角色管理列表页
	r.GET("role/list_json", RoleListJson)
	r.GET("role/add", RoleAdd)
	r.POST("role/add", RoleAddHandler)
	r.GET("role/edit", RoleEdit)
	r.POST("role/edit", RoleEditHandler)
	r.POST("role/delete", RoleDeleteHandler)

	r.GET("auth/list", AuthList)       // 权限因子列表页
	r.POST("auth/edit", AuthNodeEdit)  // 权限因子列表页
	r.POST("auth/get_nodes", GetNodes) // 权限因子列表页
	r.POST("auth/get_node", GetNode)   // 权限因子列表页
	r.POST("auth/delete", AuthDelete)  // 权限因子列表页

	r.GET("ui/icon", IconShow)
}
