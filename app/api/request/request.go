package request

type LoginForm struct {
	UserName string `json:"username" form:"username" binding:"required,min=5,max=30"`
	Password string `json:"password" form:"password" binding:"required,min=5,max=30"`
}

type LayerListForm struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

type RoleForm struct {
	LayerListForm
	ID       string `json:"id" form:"id"`
	RoleName string `json:"role_name" form:"role_name"`
	Detail   string `json:"detail" form:"detail"`
}

type AdminForm struct {
	LayerListForm
	ID        string `json:"id" form:"id"`
	LoginName string `json:"login_name" form:"login_name"`
	RealName  string `json:"real_name" form:"real_name"`
	Phone     string `json:"phone" form:"phone"`
	Email     string `json:"email" form:"email"`
}

type AdminEditForm struct {
	ID        string `json:"id" form:"id" binding:"required"`
	LoginName string `json:"login_name" form:"login_name"`
	RealName  string `json:"real_name" form:"real_name"`
	Phone     string `json:"phone" form:"phone"`
	Email     string `json:"email" form:"email"`
	Status    int
	RoleIds   string
}

type AdminAddForm struct {
	LoginName string `json:"login_name" form:"login_name"`
	Password  string `json:"password" form:"password"`
	RealName  string `json:"real_name" form:"real_name"`
	Phone     string `json:"phone" form:"phone"`
	Email     string `json:"email" form:"email"`
	Status    int
	RoleIds   string
}

type AuthNodeReq struct {
	ID       int    `json:"id" form:"id"`
	Pid      int    `json:"pid" form:"pid"`
	AuthName string `json:"auth_name" form:"auth_name"`
	AuthUrl  string `json:"auth_url" form:"auth_url"`
	Sort     int    `json:"sort" form:"sort"`
	IsShow   int    `json:"is_show" form:"is_show"`
	Icon     string `json:"icon" form:"icon"`
}

type RoleEditForm struct {
	ID        int    `json:"id" form:"id" binding:"required"`
	RoleName  string `json:"role_name" form:"role_name" binding:"required"`
	Detail    string `json:"detail" form:"detail"`
	NodesData string `json:"nodes_data" form:"nodes_data"`
}

type RoleAddForm struct {
	RoleName  string `json:"role_name" form:"role_name" binding:"required"`
	Detail    string `json:"detail" form:"detail"`
	NodesData string `json:"nodes_data" form:"nodes_data"`
}

type SiteConfForm struct {
	ID          uint   `json:"-" form:"id"`
	WebName     string `json:"web_name" form:"web_name"`
	WebUrl      string `json:"web_url" form:"web_url"`
	LogoUrl     string `json:"logo_url" form:"logo_url"`
	KeyWords    string `json:"key_words" form:"key_words"`
	Description string `json:"description" form:"description"`
	Copyright   string `json:"copyright" form:"copyright"`
	Icp         string `json:"icp" form:"icp"`
	SiteStatus  uint8  `json:"site_status"`
}

type MailConfForm struct {
	ID             uint   `json:"-" form:"id"`
	EmailName      string `json:"email_name" form:"email_name"`
	EmailHost      string `json:"email_host" form:"email_host"`
	EmailPort      string `json:"email_port" form:"email_port"`
	EmailUser      string `json:"email_user" form:"email_user"`
	EmailPwd       string `json:"email_pwd" form:"email_pwd"`
	EmailTest      string `json:"email_test" form:"email_test"`
	EmailTestTitle string `json:"email_test_title" form:"email_test_title"`
	EmailTemplate  string `json:"email_template" form:"email_template"`
	EmailStatus    int    `json:"email_status"`
}

type AvatarForm struct {
	ID     string `json:"id" form:"id" binding:"required"`
	Avatar string `json:"avatar" form:"avatar" binding:"required"`
}

type ProfileForm struct {
	RealName string `json:"real_name" form:"real_name"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
}

type PasswordForm struct {
	OldPwd     string `json:"old_pwd" form:"old_pwd" binding:"required"`
	NewPwd     string `json:"new_pwd" form:"new_pwd" binding:"required"`
	ConfirmPwd string `json:"confirm_pwd" form:"confirm_pwd" binding:"required"`
}
