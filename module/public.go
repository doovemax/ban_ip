package module

import (
	"github.com/doovemax/ban_ip/conf"
	"github.com/satyrius/gonx"
	"sync"
)

func init() {
	LogLine = make(chan string, 1000)
	IpNetfilter = make(chan string, 1000)
	IPEntry = make(chan *gonx.Entry, 1000)

	iPcountRWMutex = new(sync.RWMutex)

	//whiteIPList slice to map
	if len(conf.DefaultConfig.WhiteIPList) != 0 {
		FilterBool = true
		for _, j := range conf.DefaultConfig.WhiteIPList {
			IPListMap = map[string]bool{j: true}
		}
	}

}

var (
	LogLine     chan string
	IpNetfilter chan string
	IPEntry     chan *gonx.Entry
	FilterBool  bool
	IPListMap   map[string]bool

	//	IPconut lock
	iPcountRWMutex *sync.RWMutex
)
