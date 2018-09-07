package module

import (
	"fmt"
	"github.com/doovemax/ban_ip/conf"
	"github.com/sirupsen/logrus"
)

func IpAddFilter() (err error) {
	logrus.Println(IPcount)
	iPcountRWMutex.RLock()
	defer iPcountRWMutex.RUnlock()
	for urliptime, count := range IPcount {
		if count > conf.DefaultConfig.IPToBlackCount {
			fmt.Println(urliptime, count)
		}
	}
	return
}
