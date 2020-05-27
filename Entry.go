package logging

import (
	"fmt"
	"io"
)


/*
Format log
time [STATUS] message
 */

/*
Example:
{
	"package": "name_package"
	"function": "name_function"
	"error": "message"
}
 */

type Logs struct {
	out io.Writer
	Pkg string `json:"package"`
	Funs map[string]interface{} `json:"functions"`
	log []interface{}
	name string
}

type Fields map[string]interface{}

func LogsJson(pkg string, out io.Writer) *Logs{
	m := make(map[string]interface{})
	return &Logs{out: out, Pkg: pkg, Funs: m}
}

func (l *Logs) Log(name string) *Logs{
	l.name = name
	return &Logs{out: l.out, Pkg: l.Pkg, Funs: l.Funs,  name: l.name, log: l.log}
}

func (l *Logs) SubLogWithFields(status string, message string, time string, fields Fields) *Logs{
	f := Fields{
		"status": status,
		"message": message,
		"time": time,
		"meta": fields,
	}
	return l.updateArr(f)
}

func (l *Logs) SubLog(status string, message string, time string) *Logs{
	f := Fields{
		"status": status,
		"message": message,
		"time": time,
	}
	return l.updateArr(f)
}

func (l *Logs) updateArr(fields Fields) *Logs{
	l.log = append(l.log, fields)
	l.Funs[l.name] = l.log
	return &Logs{out: l.out, Pkg: l.Pkg, Funs: l.Funs,  name: l.name, log: l.log}
}

func (l *Logs) fieldsToStr(fields Fields) string{
	row := ""
	for key, val := range fields {
		row += fmt.Sprintf("%s: %s ", key, val)
	}
	return row
}

func (l *Logs) Report() {
	obj := structJson(l)
	fmt.Println(string(obj))
	_, _ = l.out.Write(obj)
	_, _ = l.out.Write([]byte("\n"))
}
