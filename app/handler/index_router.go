package controller

import (
	"github.com/cilidm/base-system-v2/app/middleware"
	"github.com/cilidm/base-system-v2/app/router"
)

func init() {
	r := router.New("system", "/system", middleware.AuthMiddleware)
	r.GET("/", Index) // 有问题
	r.GET("index", Index)
	r.GET("main", FrameIndex)
}
