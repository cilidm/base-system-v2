package controller

import (
	"github.com/cilidm/base-system-v2/app/middleware"
	"github.com/cilidm/base-system-v2/app/router"
)

func init() {
	r := router.New("system", "/")
	r.GET("/", middleware.CheckDefaultPage)
	r.GET("login", middleware.CheckLoginPage, Login)
	r.POST("login", LoginHandler)
	r.GET("logout", Logout)
	r.POST("isLogin", nil)
	r.GET("not_found", NotFound)
	r.GET("captcha", Captcha)
}
