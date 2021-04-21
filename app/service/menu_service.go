package service

import (
	"github.com/cilidm/base-system-v2/app/util/e"
	"github.com/cilidm/base-system-v2/app/util/gocache"
	"github.com/cilidm/toolbox/gconv"
	pkg "github.com/cilidm/toolbox/str"
	"github.com/gin-gonic/gin"

)

type SelfMenuData struct {
	ParentData []map[string]interface{}
	ChildData  []map[string]interface{}
	AllowUrl   string
}

type CacheMenu struct {
	List1    []map[string]interface{}
	List2    []map[string]interface{}
	AllowUrl string
}

func MenuService(c *gin.Context) SelfMenuData {
	var self SelfMenuData
	user := GetProfile(c)
	cheMen, found := gocache.Instance().Get(e.MenuCache + gconv.String(user.ID))
	if found && cheMen != nil { //从缓存取菜单
		menu := cheMen.(*CacheMenu)
		self.ParentData = menu.List1 //一级菜单
		self.ChildData = menu.List2  //二级菜单
		self.AllowUrl = menu.AllowUrl
	} else {
		// 左侧导航栏
		result := GetAuth(user)
		list := make([]map[string]interface{}, len(result))
		list2 := make([]map[string]interface{}, len(result))
		allowUrl := ""
		i, j := 0, 0
		for _, v := range result {
			if pkg.IsContain([]string{"", "/"}, v.AuthUrl) == false {
				if allowUrl == "" {
					allowUrl += v.AuthUrl
				} else {
					allowUrl += "," + v.AuthUrl
				}
			}
			row := make(map[string]interface{})
			if v.Pid == 1 && v.IsShow == 1 {
				row["Id"] = v.ID
				row["Sort"] = v.Sort
				row["AuthName"] = v.AuthName
				row["AuthUrl"] = v.AuthUrl
				row["Icon"] = v.Icon
				row["Pid"] = v.Pid
				list[i] = row
				i++
			}
			if v.Pid != 1 && v.IsShow == 1 {
				row["Id"] = v.ID
				row["Sort"] = v.Sort
				row["AuthName"] = v.AuthName
				row["AuthUrl"] = v.AuthUrl
				row["Icon"] = v.Icon
				row["Pid"] = v.Pid
				list2[j] = row
				j++
			}
		}
		self.ParentData = list[:i] //一级菜单
		self.ChildData = list2[:j] //二级菜单
		self.AllowUrl = allowUrl
		cheM := &CacheMenu{}
		cheM.AllowUrl = self.AllowUrl
		cheM.List1 = list[:i]
		cheM.List2 = list2[:j]
		gocache.Instance().Set(e.MenuCache+gconv.String(user.ID), cheM, 0)
	}
	return self
}

//获取用户的菜单数据
//func SelectMenuNormalByUser(userId int64) (*[]menuModel.EntityExtend, error) {
//	if userService.IsAdmin(userId) {
//		return SelectMenuNormalAll()
//	} else {
//		return SelectMenusByUserId(gconv.String(userId))
//	}
//}
