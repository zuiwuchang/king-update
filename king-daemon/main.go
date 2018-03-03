package main

import (
	"github.com/zuiwuchang/king-update/king-daemon/configure"
	"github.com/zuiwuchang/king-update/king-daemon/log"
)

const (
	//ConfigureFile 配置檔案
	ConfigureFile = "daemon.jsonnet"
)

func main() {
	//加載 配置 檔案
	if e := configure.Init(ConfigureFile); e != nil {
		log.Loggers.Fault.Fatalln(e)
	}
	log.Loggers.Info.Println(configure.GetConfigure())

	//初始化 日期
	log.Init()

}
