package controller

import (
	"github.com/cilidm/base-system-v2/app/api/request"
	"github.com/cilidm/base-system-v2/app/api/response"
	"github.com/cilidm/base-system-v2/app/model"
	"github.com/cilidm/base-system-v2/app/service"
	"github.com/cilidm/base-system-v2/app/util/e"
	"github.com/cilidm/base-system-v2/app/util/gocache"
	"github.com/cilidm/toolbox/gomail"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SiteList(c *gin.Context) {
	site, sysID := service.GetSiteConf()
	c.HTML(http.StatusOK, "sys_site_list.html", gin.H{"id": sysID, "site": site})
}

func SiteEdit(c *gin.Context) {
	var f request.SiteConfForm
	if err := c.ShouldBind(&f); err != nil {
		response.ErrorResp(c).SetType(model.OperEdit).Log(e.SiteEdit, c.Request.PostForm).WriteJsonExit()
		return
	}
	status := c.PostForm("site_status")
	if status == "on" {
		f.SiteStatus = 1
	} else {
		f.SiteStatus = 0
	}
	if err := service.SiteEditService(f); err != nil {
		response.ErrorResp(c).SetType(model.OperEdit).Log(e.SiteEdit, c.Request.PostForm).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetType(model.OperEdit).Log(e.SiteEdit, c.Request.PostForm).WriteJsonExit()
}

func MailList(c *gin.Context) {
	mail, sysID := service.GetMailConf()
	testMail := service.GetMailTestConf()
	c.HTML(http.StatusOK, "sys_mail_list.html", gin.H{"id": sysID, "mail": mail, "test": testMail})
}

func MailEdit(c *gin.Context) {
	var f request.MailConfForm
	if err := c.ShouldBind(&f); err != nil {
		response.ErrorResp(c).SetType(model.OperEdit).Log(e.MailEdit, c.Request.PostForm).WriteJsonExit()
		return
	}
	status := c.PostForm("email_status")
	if status == "on" {
		f.EmailStatus = 1
	} else {
		f.EmailStatus = 0
	}
	if err := service.MailEditService(f); err != nil {
		response.ErrorResp(c).SetType(model.OperEdit).Log(e.MailEdit, c.Request.PostForm).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetType(model.OperEdit).Log(e.MailEdit, c.Request.PostForm).WriteJsonExit()
}

func MailTest(c *gin.Context) {
	var f gomail.MailConfForm
	if err := c.ShouldBind(&f); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperEdit).Log(e.MailEdit, c.Request.PostForm).WriteJsonExit()
		return
	}
	var testMail gomail.Config
	testMail.Config = f
	testMail.MailTo = append(testMail.MailTo, f.EmailTest)
	testMail.Subject = f.EmailTestTitle
	if err := gomail.SendMail(testMail); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperEdit).Log(e.MailEdit, c.Request.PostForm).WriteJsonExit()
		return
	}
	// 保存邮件测试配置到缓存
	gocache.Instance().Set(e.TestMailConf, model.MailTest{
		EmailTest:      f.EmailTest,
		EmailTestTitle: f.EmailTestTitle,
		EmailTemplate:  f.EmailTemplate,
	}, e.TestMailEffTime)
	response.SuccessResp(c).SetType(model.OperEdit).Log(e.MailEdit, c.Request.PostForm).WriteJsonExit()
}
