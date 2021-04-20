package dao

import (
	"github.com/cilidm/base-system-v2/app/global"
	"github.com/cilidm/base-system-v2/app/model"
)

type LoginInfoDao interface {
	Insert(info model.LoginInfo) error
}

func NewLoginInfoImpl() LoginInfoDao {
	info := new(LoginInfoDaoImpl)
	return info
}

type LoginInfoDaoImpl struct {
}

func (l *LoginInfoDaoImpl) Insert(info model.LoginInfo) error {
	err := global.DBConn.Create(&info).Error
	return err
}
