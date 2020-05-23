package logging

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func openFile() (*os.File, error){
	file, err := os.OpenFile("./tmp/test.log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644)
	return file, err
}

// TestNewLog

func TestNewLog(t *testing.T) {
	fmt.Println("Start test TestNewLog")

	file, err := openFile()
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

	file, err := openFile()
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

	file, err := openFile()
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

func fun1(l *Logs) {
	l.FunsLog("fun1")
	// l.FunLog("info", "message info", time.Now().Format("2006-01-02T15:04:05"))
}

func fun2(l *Logs) {
	l.FunsLog("fun2")
}

func fun3(l *Logs) {
	l.FunsLog("fun3")
}

func TestLogger_JsonLog(t *testing.T) {
	fmt.Println("Start test TestLogger_JsonLog")

	l := LogsJson("Logging")
	fun1(l)
	fun2(l)
	fun3(l)
	l.Report()

	fmt.Println("Stop test TestLogger_JsonLog")
}