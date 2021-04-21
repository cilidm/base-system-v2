package service

import (
	"github.com/cilidm/base-system-v2/app/util/e"
	"github.com/cilidm/base-system-v2/app/util/gocache"
	"github.com/cilidm/toolbox/gconv"

	"time"
)

func Lock(loginName string) {
	gocache.Instance().Set(e.UserLock+loginName, true, time.Minute*5)
}

func UnLock(loginName string) {
	gocache.Instance().Delete(e.UserLock + loginName)
}

func CheckLock(loginName string) bool {
	res, ok := gocache.Instance().Get(e.UserLock + loginName)
	if ok && res == true {
		return true
	}
	return false
}

func SetPwdErrNum(loginName string) int {
	countNum := 0
	errNum, _ := gocache.Instance().Get(e.UserLoginErr + loginName)
	if errNum != nil {
		countNum = gconv.Int(errNum)
	}
	countNum = countNum + 1
	gocache.Instance().Set(e.UserLoginErr+loginName, countNum, time.Minute*1)
	if countNum >= 5 {
		Lock(loginName)
	}
	return countNum
}

func RemovePwdErrNum(loginName string) {
	gocache.Instance().Delete(e.UserLoginErr + loginName)
}
