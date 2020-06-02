package logging

import (
	"fmt"
	"testing"
	"time"
)

// TestLogger_JsonLog

func subfun1(l *Logs) {
	l.SubLogWithFields(
		"warning",
		"message subfun1",
		time.Now().Format("2006-01-02T15:04:05"),
		Fields{
			"fun": "subfun1",
		})
}

func fun1(l *Logs) {
	ls := l.Log("fun1")
	ls.SubLog(
		"info",
		"message fun1",
		time.Now().Format("2006-01-02T15:04:05")).Info("Info message")
	ls.SubLog(
		"warning",
		"message fun1",
		time.Now().Format("2006-01-02T15:04:05")).Warning("warning message")
	subfun1(ls)
}

func fun2(l *Logs) {
	ls := l.Log("fun2")
	ls.SubLog(
		"info",
		"message fun2",
		time.Now().Format("2006-01-02T15:04:05"),
		).Info("Info Message")
}

func fun3(l *Logs) {
	l.Log("fun3").SubLog(
		"Error",
		"message fun3",
		time.Now().Format("2006-01-02T15:04:05")).Error("Error message")
	l.Log("fun3").SubLog(
		"Info",
		"message fun3",
		time.Now().Format("2006-01-02T15:04:05")).Info("Info message")
	l.Log("fun3").SubLog(
		"other",
		"message fun3",
		time.Now().Format("2006-01-02T15:04:05")).Info("Other message")
}

func TestLogger_JsonLog(t *testing.T) {
	fmt.Println("Start test TestLogger_JsonLog")
	l := InitLog("Logging", "tmp", "json")
	fun1(l)
	fun2(l)
	fun3(l)
	l.Report()
	fmt.Println("Stop test TestLogger_JsonLog")
}

func TestLogger_SimpleLog(t *testing.T) {
	fmt.Println("Start test TestLogger_SimpleLog")
	l := InitLog("Logging", "tmp", "simple")
	fun1(l)
	fun2(l)
	fun3(l)
	l.Report()
	fmt.Println("Stop test TestLogger_SimpleLog")
}

func TestLogger_OtherLog(t *testing.T) {
	fmt.Println("Start test TestLogger_OtherLog")
	l := InitLog("Logging", "tmp", "other")
	fun1(l)
	fun2(l)
	fun3(l)
	l.Report()
	fmt.Println("Stop test TestLogger_OtherLog")
}