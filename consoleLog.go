package jlog

import (
	"fmt"
	"strings"
)

func (l *Logs) typeConsoleLog(prefix string, message string) {
	switch l.consoleLog {
	case "simple":
		l.simple(l.Pkg, l.name, prefix, message)
	case "json":
		l.json(l.Pkg, l.name, prefix, message)
	default:
		l.simple(l.Pkg, l.name, prefix, message)
	}
}

func (l *Logs) sep(status string) string{
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

func (l *Logs) pkgRow(pkg string) string{
	return pkgPrefix + ":" + brightGreen + pkg + resetColor
}

func (l *Logs) funcRow(function string) string{
	return funcPrefix + ":" + brightGreen + function + resetColor
}

func (l *Logs) detailRow(message string) string{
	return detailPrefix + ":" + message
}

func (l *Logs) simple(pkg string, function string, status string, message string) {
	row := fmt.Sprintf("%s %s %s %s", status, l.pkgRow(pkg), l.funcRow(function), l.detailRow(message))
	fmt.Println(row)
}

func (l *Logs) json(pkg string, function string, prefix string, message string) {
	status := strings.ToLower(l.fields["status"].(string))
	sep := l.sep(status)
	fmt.Printf("%s\n", sep)
	fmt.Printf("%s %s\n%s\n%s\n", prefix, message, l.pkgRow(pkg), l.funcRow(function))
	fmt.Printf("%s\n", string(structJson(l.fields)))
	fmt.Printf("%s\n", sep)
}