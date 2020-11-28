package jlog

import (
	"fmt"
	"strings"
)

const (
	pkgPrefix = brightRed + "package" + resetColor
	funcPrefix = brightRed + "func" + resetColor
	detailPrefix = magenta + "detail" + resetColor
)

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

func (l *Jlogs) sep(status string) string{
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

func (l *Jlogs) pkgRow() string{
	return pkgPrefix + ":" + brightGreen + l.Pkg + resetColor
}

func (l *Jlogs) funcRow() string{
	return funcPrefix + ":" + brightGreen + l.name + resetColor
}

func (l *Jlogs) detailRow(message string) string{
	return detailPrefix + ":" + message
}

func (l *Jlogs) simple(prefix string, message string) {
	row := fmt.Sprintf("%s %s %s %s", prefix, l.pkgRow(), l.funcRow(), l.detailRow(message))
	fmt.Println(row)
}

func (l *Jlogs) json(prefix string, message string) {
	status := strings.ToLower(l.fields["status"].(string))
	sep := l.sep(status)
	fmt.Printf("%s\n", sep)
	fmt.Printf("%s %s\n%s\n%s\n", prefix, message, l.pkgRow(), l.funcRow())
	fmt.Printf("%s\n", string(convertToJson(l.fields)))
	fmt.Printf("%s\n", sep)
}