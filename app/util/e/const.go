package e

import "time"

const (
	DefaultAvatar      = "/static/images/users/avatar-1.png"
	DefUploadSize      = 2 * 1024 * 1024 // 默认最大上传
	DefaultSaltLen     = 10
	DefaultLogPath     = "runtime/logs/"
	DefaultLogSaveName = "log"
	DefaultLogExt      = "log"

	AllowAuth = "/system,/system/index,/system/main" // 不需要验证的地址放在这里

	// 日期格式化
	TimeFormatDay  = "20060102"
	TimeFormatHour = "2006010215"
	TimeFormatAll  = "20060102150405"
	TimeFormat     = "2006-01-02 15:04:05"

	// 缓存KEY
	Menu            = "menu_list"
	AdminInfo       = "admin_info"
	Auth            = "auth"
	UserLoginErr    = "user_pwd_err_"
	UserLock        = "user_lock_"
	MaxErrNum       = 5
	MenuCache       = "menu_cache"
	TestMailConf    = "test_mail_conf"
	TestMailEffTime = time.Hour * 48

	// log title
	AdminAddHandler  = "新增管理员"
	AdminEditHandler = "修改管理员信息"
	LoginHandler     = "登陆"
	AdminDelete      = "删除管理员"

	AuthDelete   = "删除节点"
	AuthNodeEdit = "修改节点"

	RoleEditHandler   = "修改角色权限"
	RoleAddHandler    = "新增角色权限"
	RoleDeleteHandler = "删除角色权限"

	DefaultUpload = "上传图片"

	SiteEdit = "更新站点设置"
	MailEdit = "更新邮件配置"

	AvatarEdit     = "更新头像"
	ProfileEdit    = "更新资料"
	PwdEditHandler = "修改密码"

	JobServerTest = "服务器配置测试"
	JobServerDel  = "服务器配置删除"

	CronJobAddHandler = "新增定时任务"

	ImgHeight    = 80
	ImgWidth     = 240
	ImgKeyLength = 4
)
