package model

import "time"

type Auth struct {
	ID        uint
	AuthName  string
	AuthUrl   string
	UserId    int
	Pid       int
	Sort      int
	Icon      string
	IsShow    int
	Status    int
	CreateId  int
	UpdateId  int
	RoleAuths []RoleAuth `gorm:"foreignkey:AuthID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NodesResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Open bool   `json:"open"`
	Pid  int    `json:"pId"`
}

type NodeResp struct {
	ID       int    `json:"id"`
	Pid      int    `json:"pid"`
	AuthName string `json:"auth_name"`
	AuthUrl  string `json:"auth_url"`
	Sort     int    `json:"sort"`
	IsShow   int    `json:"is_show"`
	Icon     string `json:"icon"`
}
