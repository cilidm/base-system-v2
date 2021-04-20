package model

import "time"

type AdminOnline struct {
	SessionID      string    `json:"session_id" gorm:"not null default '' comment('用户会话id') VARCHAR(250)"`
	LoginName      string    `json:"login_name" gorm:"default '' comment('登录账号') VARCHAR(50)"`
	DeptName       string    `json:"dept_name" gorm:"default '' comment('部门名称') VARCHAR(50)"`
	IpAddr         string    `json:"ip_addr" gorm:"default '' comment('登录IP地址') VARCHAR(50)"`
	LoginLocation  string    `json:"login_location" gorm:"default '' comment('登录地点') VARCHAR(255)"`
	Browser        string    `json:"browser" gorm:"default '' comment('浏览器类型') VARCHAR(50)"`
	Os             string    `json:"os" gorm:"default '' comment('操作系统') VARCHAR(50)"`
	Status         string    `json:"status" gorm:"default '' comment('在线状态on_line在线off_line离线') VARCHAR(10)"`
	StartTimestamp time.Time `json:"start_timestamp" gorm:"comment('session创建时间') DATETIME"`
	LastAccessTime time.Time `json:"last_access_time" gorm:"comment('session最后访问时间') DATETIME"`
	ExpireTime     int       `json:"expire_time" gorm:"default 0 comment('超时时间，单位为分钟') INT(5)"`
}
