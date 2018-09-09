package module

import (
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
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
		// logrus.Info(":)--")

		trans <- line.Text
	}
	logrus.Fatal("tail out")
	return nil
}
