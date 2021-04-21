package main

import (
	"fmt"
	"github.com/cilidm/base-system-v2/app/core"
	"github.com/cilidm/base-system-v2/app/global"
	_ "github.com/cilidm/base-system-v2/app/handler"
	"github.com/cilidm/base-system-v2/app/router"
	"github.com/cilidm/toolbox/gconv"
	"net/http"
	"time"
)

func main() {
	core.InitConfig("./config.toml")

	global.ZapLog = core.InitZap()

	global.DBConn = core.InitConn()
	defer global.DBConn.Close()

	core.InitRedis()

	r := router.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", core.Conf.App.HttpPort),
		Handler:        r,
		ReadTimeout:    time.Duration(core.Conf.App.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(core.Conf.App.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf(`	欢迎使用 base-system
	当前版本:V2.0.1
	默认自动化文档地址:http://127.0.0.1:%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:%s
`, gconv.String(core.Conf.App.HttpPort), gconv.String(core.Conf.App.HttpPort))
	global.ZapLog.Error(s.ListenAndServe().Error())
}
