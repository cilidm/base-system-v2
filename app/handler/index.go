package controller

import (
	"github.com/cilidm/base-system-v2/app/service"
	"github.com/cilidm/base-system-v2/app/util/e"
	pkg "github.com/cilidm/toolbox/file"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"


)

func Index(c *gin.Context) {
	user := service.GetProfile(c)
	if pkg.CheckNotExist(service.GetImgSavePath(user.Avatar)) {
		user.Avatar = e.DefaultAvatar
	}
	site, _ := service.GetSiteConf()
	menus := service.MenuService(c) // AllowUrl,ParentData,ChildData
	c.HTML(http.StatusOK, "index.html", gin.H{
		"menus":     menus,
		"site":      site,
		"user":      user,
		"copyright": template.HTML(site.Copyright), // 防止转义
	})
}

func FrameIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", gin.H{})
}
