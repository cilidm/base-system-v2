package model

import "time"

type OperLog struct {
	ID            uint64
	Title         string    `json:"title" gorm:"default '' comment('模块标题') VARCHAR(50)"`
	BusinessType  int       `json:"business_type" gorm:"default 0 comment('业务类型（0其它 1新增 2修改 3删除）') INT(2)"`
	Method        string    `json:"method" gorm:"default '' comment('方法名称') VARCHAR(100)"`
	RequestMethod string    `json:"request_method" gorm:"default '' comment('请求方式') VARCHAR(10)"`
	OperatorType  int       `json:"operator_type" gorm:"default 0 comment('操作类别（0其它 1后台用户 2手机端用户）') INT(1)"`
	OperName      string    `json:"oper_name" gorm:"default '' comment('操作人员') VARCHAR(50)"`
	DeptName      string    `json:"dept_name" gorm:"default '' comment('部门名称') VARCHAR(50)"`
	OperUrl       string    `json:"oper_url" gorm:"default '' comment('请求URL') VARCHAR(255)"`
	OperIp        string    `json:"oper_ip" gorm:"default '' comment('主机地址') VARCHAR(50)"`
	OperLocation  string    `json:"oper_location" gorm:"default '' comment('操作地点') VARCHAR(255)"`
	OperParam     string    `json:"oper_param" gorm:"default '' comment('请求参数') VARCHAR(2000)"`
	JsonResult    string    `json:"json_result" gorm:"default '' comment('返回参数') VARCHAR(2000)"`
	Status        int       `json:"status" gorm:"default 0 comment('操作状态（0正常 1异常）') INT(1)"`
	ErrorMsg      string    `json:"error_msg" gorm:"default '' comment('错误消息') VARCHAR(2000)"`
	OperTime      time.Time `json:"oper_time" gorm:"comment('操作时间') DATETIME"`
}

type OperForm struct {
	Title      string
	InContent  string
	OutContent *CommonResp
}
