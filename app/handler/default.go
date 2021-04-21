package controller

import (
	"github.com/cilidm/base-system-v2/app/api/response"
	"github.com/cilidm/base-system-v2/app/core"
	"github.com/cilidm/base-system-v2/app/model"
	"github.com/cilidm/base-system-v2/app/util/e"
	f "github.com/cilidm/toolbox/file"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"time"
)

func DefaultUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperAdd).Log(e.DefaultUpload, c.Request.PostForm).WriteJsonExit()
		return
	}
	if file.Size > e.DefUploadSize {
		response.ErrorResp(c).SetMsg("文件大小超限").SetType(model.OperAdd).Log(e.DefaultUpload, file).WriteJsonExit()
		return
	}
	day := time.Now().Format(e.TimeFormatDay)
	savePath := filepath.Join(core.Conf.App.ImgSavePath, day) // 按年月日归档保存
	err = f.IsNotExistMkDir(savePath)
	if err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperAdd).Log(e.DefaultUpload, file).WriteJsonExit()
		return
	}
	if err := c.SaveUploadedFile(file, filepath.Join(savePath, file.Filename)); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperAdd).Log(e.DefaultUpload, file).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetMsg(filepath.Join(filepath.Join(core.Conf.App.ImgUrlPath, day), file.Filename)).SetType(model.OperAdd).Log(e.DefaultUpload, file).WriteJsonExit()
}


