package module

import (
	"github.com/satyrius/gonx"
	"github.com/sirupsen/logrus"
)

//日志解析
func LogParase(logformat string, input <-chan string, output chan<- *gonx.Entry) (err error) {
	var line string
	p := gonx.NewParser(logformat)
	for line = range input {
		entry, err := p.ParseString(line)
		if err != nil {
			logrus.Warn(err)
			continue
		}
		BodyBytesSent, err := entry.Field("body_bytes_sent")
		if err != nil {
			logrus.Warn(err)
			continue
		}
		// 返回字节大小为0,跳过
		if BodyBytesSent == "0" {
			// fmt.Println(BodyBytesSent)
			continue
		}
		//如果IP在白名单中,跳过
		if FilterBool {
			IP, err := entry.Field("http_x_forwarded_for")
			if err != nil {
				logrus.Warn(err)
				continue
			}
			if IPListMap[IP] == true {
				//logrus.Info(IP ," in white list")
				logrus.Debug(IP, " in white list")
				continue
			}

		}
		logrus.Info(entry)
		output <- entry
		logrus.Debug(entry)

	}
	return
}
