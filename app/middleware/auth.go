package middleware

import (
	"github.com/cilidm/base-system-v2/app/service"
	"github.com/cilidm/base-system-v2/app/util/e"
	"github.com/cilidm/base-system-v2/app/util/gocache"
	"github.com/cilidm/toolbox/gconv"
	pkg "github.com/cilidm/toolbox/str"
	"github.com/gin-gonic/gin"
	"net/http"

	"strings"
)

func AuthMiddleware(c *gin.Context) {
	if service.IsSignedIn(c) == false {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
	} else {
		if Auth(c) == false {
			c.Redirect(http.StatusFound, "/not_found")
			c.Abort()
		} else {
			if c.Request.URL.Path == "/system" {
				c.Redirect(http.StatusFound, "/system/")
			}
			c.Next()
		}
	}
}

func Auth(c *gin.Context) bool {
	url := c.Request.URL.Path
	if strings.HasSuffix(url, "/") {
		url = strings.TrimRight(url, "/")
	}
	allowAuthArr := strings.Split(e.AllowAuth, ",")
	if pkg.IsContain(allowAuthArr, url) {
		return true
	}
	var allowUrlArr []string
	user := service.GetProfile(c)
	menuCache, found := gocache.Instance().Get(e.MenuCache + gconv.String(user.ID))
	if found && menuCache != nil { //从缓存取菜单
		menu := menuCache.(*service.CacheMenu)
		allowUrlArr = strings.Split(menu.AllowUrl, ",")
	} else {
		result := service.GetAuth(user)
		for _, v := range result {
			if pkg.IsContain([]string{"", "/"}, v.AuthUrl) == false {
				allowUrlArr = append(allowUrlArr, v.AuthUrl)
			}
		}
	}
	if pkg.IsContain(allowUrlArr, url) == false {
		return false
	}
	return true
}

func CheckLoginPage(c *gin.Context) {
	if service.IsSignedIn(c) == true {
		c.Redirect(http.StatusFound, "/system/index")
	}
}

func CheckDefaultPage(c *gin.Context) {
	if service.IsSignedIn(c) == false {
		c.Redirect(http.StatusFound, "/login")
	} else {
		c.Redirect(http.StatusFound, "/system/index")
	}
}
