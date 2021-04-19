package main

import (
	"github.com/cilidm/base-system-v2/app/core"
	"github.com/cilidm/base-system-v2/app/global"
)

func main() {
	core.InitConfig("./config.toml")
	global.DBConn = core.InitConn()
	defer global.DBConn.Close()
	global.ZapLog = core.InitZap()
	core.InitRedis()
}
