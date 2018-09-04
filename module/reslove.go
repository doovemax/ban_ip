package module

import (
	"fmt"

	"github.com/satyrius/gonx"
	"github.com/sirupsen/logrus"
)

func LogParase(logformat string, trans <-chan string) error {
	var line string
	select {
	case a, ok := <-trans:
		if ok {
			line = a
		} else {
			logrus.Fatal("log to gonx chan close")
		}
	}

	p := gonx.NewParser(logformat)
	entry, err := p.ParseString(line)
	if err != nil {
		logrus.Warning(err)

	}
	fmt.Println(entry.Fields())

	return nil
}
