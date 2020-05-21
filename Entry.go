package logging

/*
Format message
{
  package: ""
  func: ""
  message ""
}
*/

/* palette
Black: \u001b[30m
Red: \u001b[31m
Green: \u001b[32m
Yellow: \u001b[33m
Blue: \u001b[34m
Magenta: \u001b[35m
Cyan: \u001b[36m
White: \u001b[37m
Reset: \u001b[0m
*/

/*
Format log
time [STATUS] message
 */

const (
	infoPrefix = "\u001B[34;1m[INFO]\u001B[0m"
	warningPrefix = "\u001B[33m[WARNING]\u001b[0m"
	errorPrefix = "\u001b[31m[ERROR]\u001b[0m"
)

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

func mapToRow(fields Fields) string{
	row := ""
	for key, val := range fields {
		row += fmt.Sprintf("%s: %s ", key, val)
	}
	return row
}

func (entry *Entry) Info(message string) {
	// entry.Fields["message"] =  message
	// fmt.Println(entry.Fields)
	// logrsl(entry.timeLog(), infoPrefix, mapToRow(entry.Fields))
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