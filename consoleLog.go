package jlog

import (
	"fmt"
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

func (l *Logs) pkgRow(pkg string) string{
	return pkgPrefix + ":" + pkg
}

func (l *Logs) funcRow(function string) string{
	return funcPrefix + ":" + function
}

func (l *Logs) detailRow(message string) string{
	return detailPrefix + ":" + message
}

func (l *Logs) simple(pkg string, function string, status string, message string) {
	row := fmt.Sprintf("%s %s %s %s", status, l.pkgRow(pkg), l.funcRow(function), l.detailRow(message))
	fmt.Println(row)
}

func (l *Logs) json(pkg string, function string, prefix string, message string) {
	fmt.Printf("%s %s\n%s\n%s\n", prefix, message, l.pkgRow(pkg), l.funcRow(function))
	fmt.Printf("%s", string(structJson(l.fields)))
	fmt.Printf("\n")
}