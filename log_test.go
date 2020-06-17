package jlog

import (
	"fmt"
	"testing"
	"time"
)

// TestLogger_JsonLog

func subfun1(l *Jlogs) {
	l.SubLogWithFields(
		"warning",
		"message subfun1",
		time.Now().Format("2006-01-02T15:04:05"),
		Fields{
			"fun": "subfun1",
		}).Warning("Warning message")
}

func fun1(l *Jlogs) {
	ls := l.Log("fun1")
	ls.SubLog(
		"info",
		"message fun1",
		time.Now().Format("2006-01-02T15:04:05")).Info("Info message")
	ls.SubLog(
		"warning",
		"message fun1",
		time.Now().Format("2006-01-02T15:04:05")).Warning("warning message")
	//ls.SubLog(
	//	"error",
	//	"message fun1",
	//	time.Now().Format("2006-01-02T15:04:05")).Error("error message")
	subfun1(ls)
}

func fun2(l *Jlogs) {
	ls := l.Log("fun2")
	ls.SubLog(
		"info",
		"message fun2",
		time.Now().Format("2006-01-02T15:04:05"),
		).Info("Info Message")
}

func fun3(l *Jlogs) {
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
	l := Init("Logging", "tmp", "json")
	fun1(l)
	fun2(l)
	fun3(l)
	fmt.Println("Stop test TestLogger_JsonLog")
}

func TestLogger_SimpleLog(t *testing.T) {
	fmt.Println("Start test TestLogger_SimpleLog")
	l := Init("Logging", "tmp", "simple")
	fun1(l)
	fun2(l)
	fun3(l)
	fmt.Println("Stop test TestLogger_SimpleLog")
}

func TestLogger_OtherLog(t *testing.T) {
	fmt.Println("Start test TestLogger_OtherLog")
	l := Init("Logging", "tmp", "other")
	fun1(l)
	fun2(l)
	fun3(l)
	fmt.Println("Stop test TestLogger_OtherLog")
}