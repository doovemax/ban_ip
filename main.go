package main

import (
	"fmt"

	"github.com/doovemax/ban_ip/conf"
	"github.com/doovemax/ban_ip/module"
)

func main() {
	fmt.Println(conf.DefaultConfig.NginxFilePath)
	fmt.Println(conf.DefaultConfig.LogFormat)
	// module.Tail(conf.DefaultConfig.NginxFilePath, module.TailConfig)
	go module.Tail(conf.DefaultConfig.NginxFilePath, module.TailConfig, module.LogLine)
	go module.LogParase(conf.DefaultConfig.LogFormat, module.LogLine)
	select {}

}
