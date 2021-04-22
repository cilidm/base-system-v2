package controller

import (
	"github.com/cilidm/base-system-v2/app/api/request"
	"github.com/cilidm/base-system-v2/app/api/response"
	"github.com/cilidm/base-system-v2/app/model"
	"github.com/cilidm/base-system-v2/app/model/dao"
	"github.com/cilidm/base-system-v2/app/service"
	"github.com/cilidm/base-system-v2/app/util/captcha"
	"github.com/cilidm/base-system-v2/app/util/e"
	"github.com/cilidm/toolbox/gconv"
	"github.com/cilidm/toolbox/ip"
	pkg "github.com/cilidm/toolbox/str"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func LoginHandler(c *gin.Context) {
	req := new(request.LoginForm)
	if err := c.ShouldBind(&req); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	isLock := service.CheckLock(req.UserName)
	if isLock {
		response.ErrorResp(c).SetMsg("账号已锁定，请稍后再试").SetType(model.OperOther).Log(e.LoginHandler, c.PostForm("username")).WriteJsonExit()
		return
	}
	userAgent := c.Request.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	ub, _ := ua.Browser()

	var info model.LoginInfo
	info.LoginName = req.UserName
	info.IpAddr = c.ClientIP()
	info.Os = ua.OS()
	info.Browser = ub
	info.LoginTime = time.Now()
	info.LoginLocation = ip.GetCityByIp(c.ClientIP())

	if sid, err := service.SignIn(req.UserName, req.Password, c); err != nil {
		errNums := service.SetPwdErrNum(req.UserName)
		having := e.MaxErrNum - errNums
		info.Msg = "账号或密码错误"
		info.Status = "0"
		err := dao.NewLoginInfoImpl().Insert(info)
		if err != nil {
			response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperOther).Log(e.LoginHandler, c.PostForm("username")).WriteJsonExit()
		}
		response.ErrorResp(c).SetMsg("账号或密码不正确,还有"+gconv.String(having)+"次之后账号将锁定").SetType(model.OperOther).Log(e.LoginHandler, c.PostForm("username")).WriteJsonExit()
	} else {
		var online model.AdminOnline
		pkg.CopyFields(&online, info)
		online.SessionID = sid
		online.Status = "on_line"
		online.ExpireTime = 1440
		online.StartTimestamp = time.Now()
		online.LastAccessTime = time.Now()
		dao.NewAdminOnlineDaoImpl().Delete(sid)
		dao.NewAdminOnlineDaoImpl().Insert(online)
		service.RemovePwdErrNum(req.UserName)

		info.Msg = "登陆成功"
		info.Status = "1"
		err := dao.NewLoginInfoImpl().Insert(info)
		if err != nil {
			response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperOther).Log(e.LoginHandler, c.PostForm("username")).WriteJsonExit()
		}
		response.SuccessResp(c).SetMsg("登陆成功").SetType(model.OperOther).Log(e.LoginHandler, c.PostForm("username")).WriteJsonExit()
	}
}

//注销
func Logout(c *gin.Context) {
	if service.IsSignedIn(c) {
		service.SignOut(c)
	}
	c.Redirect(http.StatusFound, "/login")
	c.Abort()
}

func NotFound(c *gin.Context) {
	c.HTML(http.StatusOK, "not_found.html", gin.H{})
}

func GetCaptcha(c *gin.Context) {
	id, b64s, err := captcha.CaptMake()
	if err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetData(response.CaptchaResponse{CaptchaId: id, PicPath: b64s}).WriteJsonExit()
}

func CaptchaVerify(c *gin.Context) {
	id := c.PostForm("id")
	capt := c.PostForm("capt")
	if id == "" || capt == "" {
		response.ErrorResp(c).SetMsg("param error").WriteJsonExit()
		return
	}

	if captcha.CaptVerify(id, capt) == true {
		response.SuccessResp(c).WriteJsonExit()
	} else {
		response.ErrorResp(c).SetMsg("verify failed").WriteJsonExit()
	}
	return
}
