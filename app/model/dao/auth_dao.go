package dao

import (
	"github.com/cilidm/base-system-v2/app/global"
	"github.com/cilidm/base-system-v2/app/model"
	"github.com/cilidm/toolbox/gconv"
	"strings"
)

type AuthDao interface {
	Insert(auth model.Auth) (authID uint, err error)
	Find(page, pageSize int, filters ...interface{}) (auths []*model.Auth, total int64)
	FindOne(id int) (auth model.Auth, err error)
	FindChildNode(id int) (int, error)
	Update(auth model.Auth, attr map[string]interface{}) error
	Delete(id int) error
	DeleteUse() error
}

func NewAuthDaoImpl() AuthDao {
	auth := new(AuthDaoImpl)
	return auth
}

type AuthDaoImpl struct {
}

func (a *AuthDaoImpl) DeleteUse() error {
	client := global.DBConn
	var secondAuth []model.Auth
	client.Where("pid = 20").Find(&secondAuth)
	if len(secondAuth) > 0 {
		for _, v := range secondAuth {
			client.Where("pid = ?", v.ID).Delete(model.Auth{})
			client.Where("id = ?", v.ID).Delete(model.Auth{})
		}
	}

	return nil
}

func (a *AuthDaoImpl) Insert(auth model.Auth) (authID uint, err error) {
	client := global.DBConn
	err = client.Create(&auth).Error
	return auth.ID, nil
}

func (a *AuthDaoImpl) Find(page, pageSize int, filters ...interface{}) (auths []*model.Auth, total int64) {
	offset := (page - 1) * pageSize
	client := global.DBConn
	client = client.Model(model.Auth{})
	var queryArr []string
	var values []interface{}
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			queryArr = append(queryArr, gconv.String(filters[k]))
			values = append(values, filters[k+1])
		}
	}
	client.Model(model.Auth{}).Where(strings.Join(queryArr, " AND "), values...).Order("sort,pid").Limit(pageSize).Offset(offset).Find(&auths)
	client.Model(model.Auth{}).Where(strings.Join(queryArr, " AND "), values...).Count(&total)
	return
}

func (a *AuthDaoImpl) FindOne(id int) (auth model.Auth, err error) {
	client := global.DBConn
	client.First(&auth, id)
	return auth, client.Error
}

func (a *AuthDaoImpl) FindChildNode(id int) (int, error) {
	var count int
	client := global.DBConn
	client.Model(model.Auth{}).Where("status = 1 AND pid = ?", id).Count(&count)
	return count, client.Error
}

func (a *AuthDaoImpl) Update(auth model.Auth, attr map[string]interface{}) error {
	client := global.DBConn
	client.Model(&auth).Omit("id").Updates(attr)
	return client.Error
}

func (a *AuthDaoImpl) Delete(id int) error {
	client := global.DBConn
	client.Where("id = ?", id).Delete(model.Auth{})
	return client.Error
}
