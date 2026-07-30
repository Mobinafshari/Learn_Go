package main

import "fmt"

type LogLevel int

const (
	Warn LogLevel = iota
	Error
	Trace
	Fatal
)

var levelName = []string{"Warn", "Error", "Trace", "Fatal"}

func (l LogLevel) String() string {
	if l < Warn || l > Fatal {
		return "unknown"
	}
	return levelName[l]
}

func printLog(l LogLevel) {
	fmt.Printf("%d %s\n", l, l.String())
}

func main() {
	printLog(Warn)
}
