package controller

import (
	"github.com/cilidm/base-system-v2/app/api/request"
	"github.com/cilidm/base-system-v2/app/api/response"
	"github.com/cilidm/base-system-v2/app/model"
	"github.com/cilidm/base-system-v2/app/model/dao"
	"github.com/cilidm/base-system-v2/app/service"
	"github.com/cilidm/base-system-v2/app/util/e"
	"github.com/cilidm/base-system-v2/app/util/gocache"
	"github.com/cilidm/toolbox/gconv"
	pkg "github.com/cilidm/toolbox/str"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IconShow(c *gin.Context) {
	c.HTML(http.StatusOK, "icon.html", nil)
}

func AdminAdd(c *gin.Context) {
	var rolesShow []model.RoleEditShow
	roles, err := dao.NewRoleDaoImpl().FindRoles("status = ?", "1") // 查找全部的分组
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	for _, v := range roles {
		rolesShow = append(rolesShow, model.RoleEditShow{
			ID:       gconv.Int(v.ID),
			RoleName: v.RoleName,
			Status:   v.Status,
		})
	}
	c.HTML(http.StatusOK, "admin_add.html", gin.H{"role": rolesShow})
}

func AdminAddHandler(c *gin.Context) {
	roles := c.PostFormArray("role_ids")
	roleIds := pkg.Array2Str(roles)
	status := c.PostForm("status")
	if status == "" {
		status = "0"
	}
	var f request.AdminAddForm
	if err := c.ShouldBind(&f); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperAdd).Log(e.AdminAddHandler, gin.H{"form": c.Request.PostForm}).WriteJsonExit()
		return
	}
	if err := service.AdminAddHandlerService(roleIds, status, f, c); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperAdd).Log(e.AdminAddHandler, gin.H{"form": c.Request.PostForm}).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetMsg("创建成功!").SetType(model.OperAdd).Log(e.AdminAddHandler, gin.H{"form": c.Request.PostForm}).WriteJsonExit()
	return
}

func AdminList(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_list.html", gin.H{})
}

func AdminEdit(c *gin.Context) {
	uid := c.Query("id")
	if uid == "" {
		c.String(http.StatusOK, "请检查参数")
		return
	}
	show, rolesShow, err := service.AdminEditService(uid)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	c.HTML(http.StatusOK, "admin_edit.html", gin.H{"show": show, "role": rolesShow})
}

func AdminEditHandler(c *gin.Context) {
	roles := c.PostFormArray("role_ids")
	roleIds := pkg.Array2Str(roles)
	status := c.PostForm("status")
	if status == "" {
		status = "0"
	}
	var f request.AdminEditForm
	if err := c.ShouldBind(&f); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperEdit).Log(e.AdminEditHandler, gin.H{"form": c.Request.PostForm}).WriteJsonExit()
		return
	}
	f.RoleIds = roleIds
	f.Status = gconv.Int(status)
	if err := service.UpdateAdminAttrService(f); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperEdit).Log(e.AdminEditHandler, gin.H{"form": c.Request.PostForm}).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetMsg("更新成功").SetType(model.OperEdit).Log(e.AdminEditHandler, gin.H{"form": c.Request.PostForm}).WriteJsonExit()
	return
}

func AdminListJson(c *gin.Context) {
	var f request.AdminForm
	if err := c.ShouldBind(&f); err != nil {
		response.SuccessResp(c).SetCode(0).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	count, data, err := service.AdminListJsonService(f)
	if err != nil {
		response.SuccessResp(c).SetCode(0).SetMsg(err.Error()).SetCount(count).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetCode(0).SetCount(count).SetData(data).WriteJsonExit()
}

func AdminDelete(c *gin.Context) {
	uid := c.PostForm("id")
	if uid == "" {
		response.ErrorResp(c).SetMsg("id不能为空").SetType(model.OperOther).Log(e.AdminDelete, c.Request.PostForm).WriteJsonExit()
		return
	}
	if err := service.AdminDeleteService(uid, c); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperOther).Log(e.AdminDelete, c.Request.PostForm).WriteJsonExit()
	}
	response.SuccessResp(c).SetType(model.OperOther).Log(e.AdminDelete, c.Request.PostForm).WriteJsonExit()
}

func RoleList(c *gin.Context) {
	c.HTML(http.StatusOK, "role_list.html", gin.H{})
}

func RoleListJson(c *gin.Context) {
	var f request.RoleForm
	if err := c.ShouldBind(&f); err != nil {
		response.SuccessResp(c).SetCode(0).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	count, data, err := service.RoleListJsonService(f)
	if err != nil {
		response.SuccessResp(c).SetCode(0).SetMsg(err.Error()).SetCount(count).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetCode(0).SetCount(count).SetData(data).WriteJsonExit()
}

func RoleAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "role_add.html", gin.H{})
}

func RoleAddHandler(c *gin.Context) {
	var f request.RoleAddForm
	if err := c.ShouldBind(&f); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperAdd).Log(e.RoleAddHandler, c.Request.PostForm).WriteJsonExit()
		return
	}
	if err := service.RoleAddHandlerService(f, c); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperAdd).Log(e.RoleAddHandler, c.Request.PostForm).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetType(model.OperAdd).Log(e.RoleAddHandler, c.Request.PostForm).WriteJsonExit()
}

func RoleEdit(c *gin.Context) {
	id := c.Query("id")
	role, err := service.RoleEditService(id)

	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	authID := make([]int, 0)
	for _, v := range role.RoleAuths {
		authID = append(authID, gconv.Int(v.AuthID))
	}
	c.HTML(http.StatusOK, "role_edit.html", gin.H{"role": role, "auth": authID})
}

func RoleEditHandler(c *gin.Context) {
	var f request.RoleEditForm
	if err := c.ShouldBind(&f); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperEdit).Log(e.RoleEditHandler, c.Request.PostForm).WriteJsonExit()
		return
	}
	if err := service.RoleEditHandlerService(f); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperEdit).Log(e.RoleEditHandler, c.Request.PostForm).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetType(model.OperEdit).Log(e.RoleEditHandler, c.Request.PostForm).WriteJsonExit()
}

func RoleDeleteHandler(c *gin.Context) {
	id := c.PostForm("id")
	if err := service.RoleDeleteHandlerService(id); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperOther).Log(e.RoleDeleteHandler, c.Request.PostForm).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetType(model.OperOther).Log(e.RoleDeleteHandler, c.Request.PostForm).WriteJsonExit()
}

func AuthList(c *gin.Context) {
	c.HTML(http.StatusOK, "auth_list.html", gin.H{})
}

func AuthNodeEdit(c *gin.Context) {
	var req request.AuthNodeReq
	errResp := response.ErrorResp(c).SetType(model.OperOther).Log(e.AuthNodeEdit, c.Request.PostForm)
	if err := c.ShouldBind(&req); err != nil {
		errResp.SetMsg(err.Error()).WriteJsonExit()
		return
	}
	if req.ID == 0 {
		errResp.SetType(model.OperAdd)
		if err := service.AuthInsert(req); err != nil {
			errResp.SetMsg(err.Error()).WriteJsonExit()
			return
		}
		gocache.Instance().Delete(e.MenuCache + gconv.String(service.GetUid(c))) // 删除栏目列表缓存，重新进行设置
		response.SuccessResp(c).SetType(model.OperEdit).Log(e.AuthNodeEdit, c.Request.PostForm).WriteJsonExit()
	} else {
		errResp.SetType(model.OperEdit)
		if err := service.AuthUpdate(req); err != nil {
			errResp.SetMsg(err.Error()).WriteJsonExit()
			return
		}
		gocache.Instance().Delete(e.MenuCache + gconv.String(service.GetUid(c))) // 删除栏目列表缓存，重新进行设置
		response.SuccessResp(c).SetType(model.OperEdit).Log(e.AuthNodeEdit, c.Request.PostForm).WriteJsonExit()
	}
}

func AuthDelete(c *gin.Context) {
	authID := c.PostForm("id")
	if err := service.AuthDelete(authID); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperOther).Log(e.AuthDelete, c.Request.PostForm).WriteJsonExit()
		return
	}
	gocache.Instance().Delete(e.MenuCache + gconv.String(service.GetUid(c))) // 删除栏目列表缓存，重新进行设置
	response.SuccessResp(c).SetType(model.OperOther).Log(e.AuthDelete, c.Request.PostForm).WriteJsonExit()
}

func GetNode(c *gin.Context) {
	authID := c.PostForm("id")
	resp, err := service.FindAuthByID(authID)
	if err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetData(resp).WriteJsonExit()
}

func GetNodes(c *gin.Context) {
	resp, count := service.FindAuths()
	response.SuccessResp(c).SetCount(gconv.Int(count)).SetData(resp).WriteJsonExit()
}
