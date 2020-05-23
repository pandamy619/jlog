package logging

import "fmt"


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
