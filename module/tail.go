package module

import (
	"github.com/sirupsen/logrus"

	"github.com/hpcloud/tail"
)

var (
	TailConfig tail.Config = tail.Config{
		Location:    nil,
		ReOpen:      true,
		MustExist:   true,
		Poll:        false,
		Pipe:        false,
		RateLimiter: nil,
		Follow:      true,
		MaxLineSize: 0,
		Logger:      nil,
	}
)

func Tail(configfile string, config tail.Config, trans chan<- string) error {
	t, err := tail.TailFile(configfile, config)
	if err != nil {
		return err
	}
	for line := range t.Lines {
		logrus.Info("[", line.Text, "]")

		trans <- line.Text
	}
	return nil
}
