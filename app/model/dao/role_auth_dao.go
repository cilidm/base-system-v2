package dao

import (
	"bytes"
	"github.com/cilidm/base-system-v2/app/global"
	"github.com/cilidm/base-system-v2/app/model"
	"github.com/cilidm/toolbox/gconv"
	"strings"
)

type RoleAuthDao interface {
	GetByRoleIds(RolesIds string) (authIds string, err error)
	Insert(auth model.RoleAuth) error
	InsertByTx(auths []model.RoleAuth) error
	Delete(k, v string) error
}

func NewRoleAuthDaoImpl() RoleAuthDao {
	roleAuth := new(RoleAuthDaoImpl)
	return roleAuth
}

type RoleAuthDaoImpl struct {
	//rw sync.RWMutex
}

func (r *RoleAuthDaoImpl) GetByRoleIds(RolesIds string) (authIds string, err error) {
	ids := strings.Split(RolesIds, ",")
	var roleAuths []model.RoleAuth
	err = global.DBConn.Model(model.RoleAuth{}).Where("role_id IN (?)", ids).Find(&roleAuths).Error
	if err != nil {
		return "", err
	}
	b := bytes.Buffer{}
	for _, v := range roleAuths {
		if v.AuthID != 0 && v.AuthID != 1 {
			b.WriteString(gconv.String(v.AuthID))
			b.WriteString(",")
		}
	}
	authIds = strings.TrimRight(b.String(), ",")
	return authIds, nil
}

func (r *RoleAuthDaoImpl) InsertByTx(auths []model.RoleAuth) error {
	for _, v := range auths {
		err := global.DBConn.Create(&v).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RoleAuthDaoImpl) Insert(auth model.RoleAuth) error {
	err := global.DBConn.Create(&auth).Error
	return err
}

func (r *RoleAuthDaoImpl) Delete(k, v string) error {
	tx := global.DBConn.Begin()
	err := tx.Where(k, v).Delete(model.RoleAuth{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
