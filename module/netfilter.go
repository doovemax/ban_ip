package module

import (
	"github.com/doovemax/ban_ip/conf"
	"time"
)

func ipAddFiter() (err error) {
	time.Sleep(time.Second * 30)
	for urlip, count := range IPcount {
		if count > conf.DefaultConfig.IPToBlackCount {

		}
	}
	return
}
