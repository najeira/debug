package debug

import (
	"fmt"
	"log"
	"os"
)

var (
	Logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
)

func Print(v interface{}, args ...interface{}) {
	if Logger == nil {
		return
	}
	if len(args) > 0 {
		if format, ok := v.(string); ok {
			Logger.Output(2, fmt.Sprintf(format, args...))
			return
		}
	}
	Logger.Output(2, fmt.Sprint(v))
}
