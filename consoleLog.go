package jlog

import (
	"fmt"
	"time"
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

func (l *Logs) simple(pkg string, function string, status string, message string) {
	datetime := time.Now().Format("2006-01-02T15:04:05")
	fmt.Printf("%s %s \u001b[32mpackage\u001B[0m:%s \u001B[33;1mfunc\u001B[0m:%s %s\n",
		status,
		datetime,
		pkg,
		function,
		message,
		)
}

func (l *Logs) json(pkg string, function string, prefix string, message string) {
	fmt.Printf("%s %s\n", prefix, message)
	fmt.Printf("\u001B[32m[package]\u001B[0m %s \u001B[33;1m[func]\u001B[0m %s\n", pkg, function)
	fmt.Printf("%s", string(structJson(l.fields)))
	fmt.Printf("\n")
}