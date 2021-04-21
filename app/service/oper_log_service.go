package service

import (
	"encoding/json"
	"errors"
	"github.com/cilidm/base-system-v2/app/model"
	"github.com/cilidm/base-system-v2/app/model/dao"
	"github.com/cilidm/toolbox/ip"
	"github.com/gin-gonic/gin"

	"time"
)

func CreateOperLog(c *gin.Context, f model.OperForm) error {
	user := GetProfile(c)
	if user == nil {
		return errors.New("用户未登陆")
	}
	outJson, _ := json.Marshal(f.OutContent)
	var oper model.OperLog
	oper.Title = f.Title
	oper.OperParam = f.InContent
	oper.JsonResult = string(outJson)
	oper.BusinessType = int(f.OutContent.Type)
	oper.OperatorType = 1 // 操作类别（0其它 1后台用户 2手机端用户）
	if f.OutContent.Code == 0 {
		oper.Status = 0
	} else {
		oper.Status = 1
	}
	oper.OperName = user.LoginName
	oper.RequestMethod = c.Request.Method
	oper.OperUrl = c.Request.URL.Path
	oper.Method = c.Request.Method
	oper.OperIp = c.ClientIP()
	oper.OperLocation = ip.GetCityByIp(oper.OperIp)
	oper.OperTime = time.Now()
	if err := dao.NewOperLogDaoImpl().Insert(oper); err != nil {
		return err
	}
	return nil
}
