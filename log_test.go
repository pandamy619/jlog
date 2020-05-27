package logging

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func openFile(filePath string) (*os.File, error){
	file, err := os.OpenFile(filePath,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644)
	return file, err
}

// TestNewLog

func TestNewLog(t *testing.T) {
	fmt.Println("Start test TestNewLog")

	file, err := openFile("./tmp/test.log")
	defer file.Close()
	if err != nil {
		t.Error("file opening error")
		os.Exit(1)
	}
	NewLog(time.Now().Format("2006-01-02T15:04:05"), file)

	fmt.Println("Stop test TestNewLog")
}

// TestLogger_Output

func TestLogger_Output(t *testing.T) {
	fmt.Println("Start TestLogger_Output")

	file, err := openFile("./tmp/test.log")
	defer file.Close()
	if err != nil {
		t.Error("file opening error")
		os.Exit(1)
	}
	l := NewLog(time.Now().Format("2006-01-02T15:04:05"), file)

	err = l.Output("test row 1")
	if err != nil {
		t.Error("error write to file")
	}
	err = l.Output("test row 2")
	fmt.Println("Stop TestLogger_Output")
}

// TestLogger_Print

func TestLogger_Print(t *testing.T) {
	fmt.Println("Start test TestLogger_Print")

	file, err := openFile("./tmp/test.log")
	defer file.Close()
	if err != nil {
		t.Error("file opening error")
		os.Exit(1)
	}

	l := NewLog(time.Now().Format("2006-01-02T15:04:05"), file)
	l.WithFields(Fields{
		"package": "logging",
		"function": "TestLogger_Info",
	})
	l.Info("message info")
	l.Warning("message warning")
	l.Error("message error")
	fmt.Println("Stop test TestLogger_Print")
}

// TestLogger_JsonLog

func subfun1(l *Logs) {
	l.SubLogWithFields(
		"warning",
		"message subfun1",
		time.Now().Format("2006-01-02T15:04:05"),
		Fields{"fun": "subfun1"})
}

func fun1(l *Logs) {
	ls := l.Log("fun1").SubLog(
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
	ls.SubLog("warning", "message fun2", time.Now().Format("2006-01-02T15:04:05"))
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
}

func TestLogger_JsonLog(t *testing.T) {
	fmt.Println("Start test TestLogger_JsonLog")
	file, err := openFile("./tmp/test.json")
	defer file.Close()
	if err != nil {
		t.Error("file opening error")
	}
	l := LogsJson("Logging", file)
	fun1(l)
	fun2(l)
	fun3(l)
	defer l.Report()
	fmt.Println("Stop test TestLogger_JsonLog")
}