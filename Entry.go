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

/*
func (entry *Entry) timeLog() string{
	// t := time.Now()
	if entry.FormatTime == "" {
		return entry.Time.String()
	}
	return entry.Time.Format(entry.FormatTime)
}

func logrsl(time string, prefix string, row string) {
	message := fmt.Sprintf("%s %s %s", time, prefix, row)
	fmt.Println(message)
}

func (entry *Entry) Warning(message string) {
	// entry.Fields["message"] =  message
	// logrsl(entry.timeLog(), warningPrefix, mapToRow(entry.Fields))
}

func (entry *Entry) Error(message string) {
	// entry.Fields["message"] =  message
	// logrsl(entry.timeLog(), errorPrefix, mapToRow(entry.Fields))
}

 */

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
