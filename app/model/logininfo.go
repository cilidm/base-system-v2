package model

import (
	"time"
)

type LoginInfo struct {
	InfoId        int64     `json:"info_id" gorm:"primary_key"`
	LoginName     string    `json:"login_name" gorm:"default '' comment('登录账号') VARCHAR(50)"`
	IpAddr        string    `json:"ip_addr" gorm:"default '' comment('登录IP地址') VARCHAR(50)"`
	LoginLocation string    `json:"login_location" gorm:"default '' comment('登录地点') VARCHAR(255)"`
	Browser       string    `json:"browser" gorm:"default '' comment('浏览器类型') VARCHAR(50)"`
	Os            string    `json:"os" gorm:"default '' comment('操作系统') VARCHAR(50)"`
	Status        string    `json:"status" gorm:"default '0' comment('登录状态（0成功 1失败）') CHAR(1)"`
	Msg           string    `json:"msg" gorm:"default '' comment('提示消息') VARCHAR(255)"`
	LoginTime     time.Time `json:"login_time" gorm:"comment('访问时间') DATETIME"`
}
