package module

import (
	"fmt"

	"github.com/satyrius/gonx"
	"github.com/sirupsen/logrus"
)

func LogParase(logformat string, trans <-chan string) (err error) {
	var line string
	p := gonx.NewParser(logformat)
	for line = range trans {
		entry, err := p.ParseString(line)
		if err != nil {
			logrus.Warning(err)
			return
		}
		fmt.Println(entry.Fields())
	}
	return
}
