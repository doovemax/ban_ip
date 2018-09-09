package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/doovemax/ban_ip/conf"
	"github.com/doovemax/ban_ip/module"
)

func main() {
	fmt.Println(conf.DefaultConfig.NginxFilePath)
	fmt.Println(conf.DefaultConfig.LogFormat)
	fmt.Println(conf.DefaultConfig.WhiteIPList)

	// module.Tail(conf.DefaultConfig.NginxFilePath, module.TailConfig)
	logrus.Println("进入tail")
	time.Sleep(time.Second * 3)
	go module.Tail(conf.DefaultConfig.NginxFilePath, module.TailConfig, module.LogLine)
	logrus.Println("进入parase")
	time.Sleep(time.Second * 3)
	go module.LogParase(conf.DefaultConfig.LogFormat, module.LogLine, module.IPEntry)
	logrus.Println("进入count")
	time.Sleep(time.Second * 3)
	go module.Count(module.IPEntry)
	logrus.Println("进入filter")
	time.Sleep(time.Second * 3)
	module.CronWork()
	select {}

}
