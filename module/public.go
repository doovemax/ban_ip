package module

func init() {
	LogLine = make(chan string, 1000)
}

var (
	LogLine chan string
)
