package jlog

import (
	"fmt"
	"strings"
)

const (
	prefixPkg    = brightRed + "package" + resetColor
	prefixFunc   = brightRed + "func" + resetColor
	prefixDetail = magenta + "detail" + resetColor
)

// TODO поменять наименование метода.
// Вынести логику определения типа вывода в консоль в адаптер
func (l *Jlogs) typeConsoleLog(prefix string, message string) {
	switch l.consoleType {
	case "simple":
		l.simple(prefix, message)
	case "json":
		l.json(prefix, message)
	default:
		l.simple(prefix, message)
	}
}

func (l *Jlogs) sep(status string) string {
	var color string
	switch status {
	case "info":
		color = infoColor
	case "warning":
		color = warningColor
	case "error":
		color = errorColor
	default:
		color = infoColor
	}
	return fmt.Sprintf("%s-------------------------%s", color, resetColor)
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
		fmt.Sprintf("%s %s %s %s", prefix, l.getPkgName(), l.getFuncName(), l.addPrefix(prefixDetail, message)),
	)
}

func (l *Jlogs) json(prefix string, message string) {
	status := strings.ToLower(l.fields["status"].(string))
	sep := l.sep(status)
	fmt.Printf("%s\n", sep)
	fmt.Printf("%s %s\n%s\n%s\n", prefix, message, l.getPkgName(), l.getFuncName())
	fmt.Printf("%s\n", string(convertToJson(l.fields)))
	fmt.Printf("%s\n", sep)
}
