package module

import (
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

func init() {
	cronWorker = cron.New()
}


var (
	cronWorker  *cron.Cron
)


func CronWork (){
	cronWorker.AddFunc("*/30 * * * * *", func() {
		logrus.Info("ip add filter go run")
		err := IpAddFilter()
		if err !=nil {
			logrus.Warn(err)
		}
	})




	cronWorker.Start()
}