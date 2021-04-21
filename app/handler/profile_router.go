package controller

import (
	"github.com/cilidm/base-system-v2/app/middleware"
	"github.com/cilidm/base-system-v2/app/router"
)

func init() {
	r := router.New("system", "/system", middleware.AuthMiddleware)
	r.GET("user/edit", Profile)
	r.POST("user/edit", ProfileEdit)
	r.POST("user/avatar", AvatarEdit)

	r.GET("user/pwd", PwdEdit)
	r.POST("user/pwd", PwdEditHandler)
}
