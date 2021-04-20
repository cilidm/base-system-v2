package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Admin struct {
	ID        uint      `json:"id"`
	LoginName string    `json:"login_name" gorm:"unique_index"`
	RealName  string    `json:"real_name"`
	Password  string    `json:"-"`
	Level     int       `json:"level"`
	RoleIds   string    `json:"role_ids"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	Salt      string    `json:"-"`
	LastIp    string    `json:"-"`
	LastLogin time.Time `json:"last_login"`
	Status    int       `json:"status"`
	CreateId  int       `json:"create_id"`
	UpdateId  int       `json:"update_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-" gorm:"default:null"`
}

func (a *Admin) BeforeCreate(scope *gorm.Scope) {
}

func (a *Admin) BeforeUpdate(scope *gorm.Scope) {
}

type AdminShow struct {
	ID        uint   `json:"id"`
	LoginName string `json:"login_name"`
	RealName  string `json:"real_name"`
	RoleIds   string `json:"role_ids"`
	Level     int    `json:"level"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Status    int    `json:"status"`
	LastLogin string `json:"last_login"`
	LastIp    string `json:"last_ip"`
}

type RoleShow struct {
	ID        int    `json:"id"`
	RoleName  string `json:"role_name"`
	Detail    string `json:"detail"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
