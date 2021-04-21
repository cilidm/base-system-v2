package service

import (
	"errors"
	"github.com/cilidm/base-system-v2/app/api/request"
	"github.com/cilidm/base-system-v2/app/model"
	"github.com/cilidm/base-system-v2/app/model/dao"
	"github.com/cilidm/toolbox/gconv"
	pkg "github.com/cilidm/toolbox/str"

)

func AuthInsert(req request.AuthNodeReq) (err error) {
	var auth model.Auth
	auth.Pid = req.Pid
	auth.AuthName = req.AuthName
	auth.AuthUrl = req.AuthUrl
	auth.Icon = req.Icon
	auth.Sort = req.Sort
	auth.IsShow = req.IsShow
	auth.Status = 1
	_, err = dao.NewAuthDaoImpl().Insert(auth)
	return err
}

func AuthUpdate(req request.AuthNodeReq) (err error) {
	auth, err := dao.NewAuthDaoImpl().FindOne(req.ID)
	if err != nil {
		return err
	}
	err = dao.NewAuthDaoImpl().Update(auth, pkg.Struct2MapByTag(req, "json"))
	return err
}

func AuthDelete(authID string) error {
	count, err := dao.NewAuthDaoImpl().FindChildNode(gconv.Int(authID))
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("请先删除子节点")
	}
	err = dao.NewRoleAuthDaoImpl().Delete("auth_id = ?", authID)
	if err != nil {
		return err
	}
	err = dao.NewAuthDaoImpl().Delete(gconv.Int(authID))
	return err
}

func FindAuthByID(authID string) (*model.NodeResp, error) {
	auth, err := dao.NewAuthDaoImpl().FindOne(gconv.Int(authID))
	if err != nil {
		return nil, err
	}
	return &model.NodeResp{
		ID:       gconv.Int(auth.ID),
		Pid:      auth.Pid,
		AuthName: auth.AuthName,
		AuthUrl:  auth.AuthUrl,
		Sort:     auth.Sort,
		IsShow:   auth.IsShow,
		Icon:     auth.Icon}, nil
}

func FindAuths() ([]model.NodesResp, int64) {
	filters := make([]interface{}, 0)
	filters = append(filters, "status = ?", 1)
	auths, count := dao.NewAuthDaoImpl().Find(1, 5000, filters...)
	var resp []model.NodesResp
	for _, v := range auths {
		resp = append(resp, model.NodesResp{
			ID:   gconv.Int(v.ID),
			Name: v.AuthName,
			Open: true,
			Pid:  v.Pid,
		})
	}
	return resp, count
}
