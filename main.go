package main

import (
	"fmt"
	"time"

	"github.com/doovemax/ban_ip/conf"
	"github.com/doovemax/ban_ip/module"
)

func main() {
	fmt.Println(conf.DefaultConfig.NginxFilePath)
	fmt.Println(conf.DefaultConfig.LogFormat)
	fmt.Println(len(conf.DefaultConfig.WhiteIPList))

	// module.Tail(conf.DefaultConfig.NginxFilePath, module.TailConfig)
	time.Sleep(time.Second * 10)
	go module.Tail(conf.DefaultConfig.NginxFilePath, module.TailConfig, module.LogLine)
	go module.LogParase(conf.DefaultConfig.LogFormat, module.LogLine, module.IPEntry)
	go module.Count(module.IPEntry)
	select {}

}
