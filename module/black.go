package module

import (
	"github.com/satyrius/gonx"
	"github.com/sirupsen/logrus"
)

func init() {
	IPcount = make(map[string]int)
}

var (
	//IPcount
	IPcount map[string]int
)

func Count(inEntry <-chan *gonx.Entry) (err error) {
	for entry := range inEntry {
		Uri, err := entry.Field("request_uri")
		if err != nil {
			logrus.Warn(err)
			continue
		}
		IP, err := entry.Field("http_x_forwarded_for")
		if err != nil {
			logrus.Warn(err)
			continue
		}
		//if  _,ok := IPcount[Uri+":"+IP]; !ok {
		//	IPcount = map[string]int{Uri+":"+IP:1}
		//} else {
		//	IPcount[Uri+":"+IP]++
		//}
		IPcount[Uri+":"+IP]++

		//f,_:=jsoniter.MarshalIndent(IPcount,"","    ")
		//fmt.Println(string(f))
		//os.Exit(0)
	}
	return
}
