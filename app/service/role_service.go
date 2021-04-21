package service

import (
	"errors"
	"github.com/cilidm/base-system-v2/app/api/request"
	"github.com/cilidm/base-system-v2/app/model"
	"github.com/cilidm/base-system-v2/app/model/dao"
	"github.com/cilidm/base-system-v2/app/util/e"
	"github.com/cilidm/toolbox/gconv"
	pkg "github.com/cilidm/toolbox/str"
	"github.com/gin-gonic/gin"

	"strings"
)

func BuildRoleFilter(r request.RoleForm) []interface{} {
	filters := make([]interface{}, 0)
	if r.ID != "" {
		filters = append(filters, "id = ?", r.ID)
	}
	if r.RoleName != "" {
		filters = append(filters, "role_name LIKE ?", "%"+r.RoleName+"%")
	}
	if r.Detail != "" {
		filters = append(filters, "detail LIKE ?", "%"+r.Detail+"%")
	}
	return filters
}

func RoleListJsonService(f request.RoleForm) (count int, data []map[string]interface{}, err error) {
	if f.Page == 0 {
		f.Page = 1
	}
	if f.Limit == 0 {
		f.Limit = 10
	}
	filters := BuildRoleFilter(f)
	list, count, err := dao.NewRoleDaoImpl().FindByPage(f.Page, f.Limit, filters...)
	if err != nil {
		return count, data, err
	}
	for _, v := range list {
		data = append(data, pkg.Struct2MapByTag(model.RoleShow{
			ID:        gconv.Int(v.ID),
			RoleName:  v.RoleName,
			Detail:    v.Detail,
			CreatedAt: v.CreatedAt.Format(e.TimeFormat),
			UpdatedAt: v.UpdatedAt.Format(e.TimeFormat),
		}, "json"))
	}
	return count, data, nil
}

func RoleEditService(roleID string) (model.Role, error) {
	role, err := dao.NewRoleDaoImpl().FindOne(roleID)
	return role, err
}

func RoleAddHandlerService(f request.RoleAddForm, c *gin.Context) error {
	hasRole, err := dao.NewRoleDaoImpl().FindRole("role_name = ?", f.RoleName)
	if err == nil || hasRole.ID > 0 {
		return errors.New("角色名已存在")
	}
	roleID, err := dao.NewRoleDaoImpl().Insert(model.Role{
		RoleName: f.RoleName,
		Detail:   f.Detail,
		CreateId: GetUid(c),
	})
	if err != nil {
		return err
	}
	if strings.HasSuffix(f.NodesData, ",") {
		f.NodesData = string([]rune(f.NodesData)[:len(f.NodesData)-1])
	}
	nodesArr := strings.Split(f.NodesData, ",")
	for _, v := range nodesArr {
		err := dao.NewRoleAuthDaoImpl().Insert(model.RoleAuth{
			RoleID: gconv.Uint64(roleID),
			AuthID: gconv.Uint(v),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func RoleEditHandlerService(f request.RoleEditForm) error {
	role, err := dao.NewRoleDaoImpl().FindOne(gconv.String(f.ID))
	if err != nil {
		return err
	}
	attr := make(map[string]interface{})
	attr["role_name"] = f.RoleName
	attr["detail"] = f.Detail
	if err := dao.NewRoleDaoImpl().Update(role, attr); err != nil {
		return err
	}
	if err := dao.NewRoleAuthDaoImpl().Delete("role_id = ?", gconv.String(f.ID)); err != nil {
		return err
	}
	if strings.HasSuffix(f.NodesData, ",") {
		f.NodesData = string([]rune(f.NodesData)[:len(f.NodesData)-1])
	}
	authIds := strings.Split(f.NodesData, ",")
	var roleAuths []model.RoleAuth
	for _, v := range authIds {
		var ra model.RoleAuth
		ra.RoleID = gconv.Uint64(f.ID)
		ra.AuthID = gconv.Uint(v)
		roleAuths = append(roleAuths, ra)
	}
	if err := dao.NewRoleAuthDaoImpl().InsertByTx(roleAuths); err != nil {
		return err
	}
	return nil
}

func RoleDeleteHandlerService(id string) error {
	role, err := dao.NewRoleDaoImpl().FindOne(id)
	if err != nil {
		return err
	}
	err = dao.NewRoleDaoImpl().Delete(role)
	if err != nil {
		return err
	}
	err = dao.NewRoleAuthDaoImpl().Delete("role_id = ?", gconv.String(role.ID))
	return err
}
