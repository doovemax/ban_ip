package module

import (
	"fmt"
	"github.com/doovemax/ban_ip/conf"
	"time"
)

func IpAddFilter() (err error) {
	time.Sleep(time.Second * 30)
	for urlip, count := range IPcount {
		if count > conf.DefaultConfig.IPToBlackCount {
			fmt.Println(urlip, count)
		}
	}
	return
}
