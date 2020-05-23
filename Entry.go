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

type Field map[string]interface{}

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

// попытка строгого назначения полей
type Log struct {
	stauts string
	message string
	time string
	// TODO Возможность пользователю добавлять кастомные поля
	// Fields
}

type Funs struct {
	Log map[string]interface{}
}

type Logs struct {
	Pkg string `json:"package"`
	Funs []interface{} `json:"functions"`
}

func (l *Logs) FunLog(status string, message string, time string) {
	l.Funs = append(l.Funs, Fields{
		"status": status,
		"message": message,
		"time": time,
	})
}

func (l *Logs) FunsLog(name string) {
	l.Funs = append(l.Funs, Fields{
		name: []string{},
	})
}

func LogsJson(pkg string) *Logs{
	return &Logs{Pkg: pkg}
}

func (l *Logs) Report()  {
	obj := structJson(l)
	fmt.Println(string(obj))
}
