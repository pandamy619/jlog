package logging

import (
	"fmt"
	"time"
)

func (l *Logs) consoleLog(pkg string, function string, status string, message string) {
	datetime := time.Now().Format("2006-01-02T15:04:05")
	fmt.Printf("%s \u001b[32mPACKAGE\u001B[0m: %s \u001B[33;1mfunction\u001B[0m: %s %s %s\n",
		datetime,
		pkg,
		function,
		status,
		message,
		)
	// l.outFile()
}
