package debug

import (
	"fmt"
	"log"
	"os"
)

var (
	Logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
)

func Printf(format string, args ...interface{}) {
	if Logger == nil {
		return
	}
	Logger.Output(2, fmt.Sprintf(format, args...))
}

func Print(args ...interface{}) {
	if Logger == nil {
		return
	}
	Logger.Output(2, fmt.Sprint(args...))
}
