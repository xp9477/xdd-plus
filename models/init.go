package models

import (
	// "fmt"
	"os"
	"path/filepath"
	// "strconv"
	// "strings"

	// "github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
)

var test2 = func(string) {

}

func init() {
	killp()
	for _, arg := range os.Args {
		if arg == "-d" {
			Daemon()
		}
	}
	ExecPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	logs.Info("当前%s", ExecPath)
	initConfig()
	// initNolan()
	initDB()
	go initVersion()
	// go initUserAgent()
	initContainer()
	initHandle()
	initCron()
	go initTgBot()
	InitReplies()
	initTask()
	initNolan()
	// initRepos()
	intiSky()
}

func initNolan() {
	// a := GetLocalMac()
	//
	// logs.Info(fmt.Printf("您的设备码是:%s,请发送给管理员进行认证\n", a))

	// s, _ := httplib.Get(fmt.Sprintf("http://auth.smxy.xyz/user/auth?qqNum=%s", strconv.FormatInt(Config.QQVID, 10))).String()
	// contains := strings.Contains(s, "true")
	// if contains {
	// 	Config.VIP = true
	// 	logs.Info("VIP验证成功")
	// } else {
	// 	logs.Info("VIP校验失败")
	// }
	Config.VIP = true
}
