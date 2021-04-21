package controller

import (
	"github.com/cilidm/base-system-v2/app/middleware"
	"github.com/cilidm/base-system-v2/app/router"
)

func init() {
	r := router.New("system", "/system", middleware.AuthMiddleware)
	r.POST("upload/def_upload", DefaultUpload)
}
