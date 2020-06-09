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

func (l *Logs) simple(pkg string, function string, status string, message string) {
	fmt.Printf("%s \u001b[32mpackage\u001B[0m:%s \u001B[33;1mfunc\u001B[0m:%s detail:%s\n",
		status,
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