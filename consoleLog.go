package jlog

import (
	"fmt"
)

const (
	prefixPkg    = brightRed + "package" + resetColor
	prefixFunc   = brightRed + "func" + resetColor
	prefixDetail = magenta + "detail" + resetColor
)

func (l *Jlogs) consoleLog(prefix string, message string) {
	l.switchConsoleLog(l.consoleType, prefix, message)

}

func (l *Jlogs) switchConsoleLog(typeLog string, prefix string, message string) {
	switch typeLog {
	case "log":
		l.simple(prefix, message)
	// case "json":
	//	l.json(prefix, message)
	default:
		// TODO добавить оповещение, что метод не реализован
		l.simple(prefix, message)
	}
}

// getPkg return colorized name package.
func (l *Jlogs) getPkgName() string {
	return prefixPkg + ":" + brightGreen + l.Pkg + resetColor
}

// getFunc return colorized name function.
func (l *Jlogs) getFuncName() string {
	return prefixFunc + ":" + brightGreen + l.name + resetColor
}

// addPrefix message then add colorized prefix 'detail' to the message and return it.
func (l *Jlogs) addPrefix(prefix string, message string) string {
	return prefix + ":" + message
}

// simple calls Println to print row to the console.
// Row is formed from status/name package/name function/message.
func (l *Jlogs) simple(prefix string, message string) {
	fmt.Println(
		fmt.Sprintf(
			"%s %s %s %s", prefix, l.getPkgName(), l.getFuncName(), l.addPrefix(prefixDetail, message),
		),
	)
}
