package logging

import (
	"fmt"
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

type Fields map[string]interface{}

// type Field map[string]interface{}

func (l *Logger) fieldsToStr(fields Fields) string{
	row := ""
	for key, val := range fields {
		row += fmt.Sprintf("%s: %s ", key, val)
	}
	return row
}

func (l *Logger) WithFields(fields Fields) {
	l.fields = l.fieldsToStr(fields)
}


type Logs struct {
	Pkg string `json:"package"`
	Funs map[string]interface{} `json:"functions"`
	log []interface{}
	name string
}

func (l *Logs) SubLog(status string, message string, time string) *Logs{
	// TODO Возможность пользователю добавлять кастомные поля
	// FunLog(status string, message string, time string, fields Fields)
	// fields: Словарик для кастомного назначения полей
	l.log = append(l.log, Fields{
		"status": status,
		"message": message,
		"time": time,
	})
	l.Funs[l.name] = l.log
	return &Logs{Pkg: l.Pkg, Funs: l.Funs,  name: l.name, log: l.log}
}

func (l *Logs) Log(name string) *Logs{
	l.name = name
	return &Logs{Pkg: l.Pkg, Funs: l.Funs,  name: l.name, log: l.log}
}

func LogsJson(pkg string) *Logs{
	m := make(map[string]interface{})
	return &Logs{Pkg: pkg, Funs: m}
}

func (l *Logs) Report()  {
	obj := structJson(l)
	fmt.Println(string(obj))
}
