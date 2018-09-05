package module

import (
	"github.com/satyrius/gonx"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
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
		aTime, err := entry.Field("time_local")
		if err != nil {
			logrus.Warn(err)
			continue
		}
		aTimeFormat, err := time.Parse("2/Jan/2006:15:04:05 -0700", aTime)
		if err != nil {
			logrus.Warn(err)
			continue
		}
		aTimeMin := strconv.Itoa(aTimeFormat.Year()) + strconv.Itoa(int(aTimeFormat.Month())) + strconv.Itoa(aTimeFormat.Day()) + strconv.Itoa(aTimeFormat.Hour()) + strconv.Itoa(aTimeFormat.Minute())

		IPcount[Uri+":"+IP+":"+aTimeMin]++

		//f,_:=jsoniter.MarshalIndent(IPcount,"","    ")
		//fmt.Println(string(f))
		//os.Exit(0)
	}
	return
}
