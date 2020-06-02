# logger

## Installation
```bash
go get github.com/pandamy619/logger
```

## Usage

### Init struct

```go
func InitLog(pkg string, location string, consoleLog string) *Logs
```
* pkg: Name package.
* location: Folder where logs will be saved.
* consoleLog: Method output to console, simple or json.

#### Example

```go
package main
import (
    "logging"
)

func main() {
   l := InitLog("SomePackage", "tmp", "simple")
}
```

### log usage

```go
func Log(name string) *Logs
```
* name: Function name.

#### Example

```go
package main
import (
    "logging"
)

func main() {
   l := logging.InitLog("SomePackage", "tmp", "simple")
   l.Log("SomeNameFunction")
}
```

### Sublog usage

```go
func SubLog(status string, message string, time string) *Logs
```
* status: Status log (info/warning/error/customStatus).
* message: Message containing information about the log.
* time: Date and time.

#### Example

```go
package main
import (
    "time"

    "logging"
)

func main() {
   l := logging.InitLog("SomePackage", "tmp", "simple")
   l.Log("SomeNameFunction").SubLog(
       "error",
       "some message",
       time.Now().Format("2006-01-02T15:04:05"),
   )
}
```
### SubLog with custom field
```go
func SubLogWithFields(status string, message string, time string, field Fields) *Logs
```
* status: Status log (info/warning/error/customStatus).
* message: Message containing information about the log.
* time: Date and time.
* fields: Fields.

#### Example

```go
package main
import (
    "time"

    "logging"
)

func main() {
   l := logging.InitLog("SomePackage", "tmp", "simple")
   l.Log("SomeNameFunction").SubLogWithFields(
       "error",
       "some message",
       time.Now().Format("2006-01-02T15:04:05"),
       logging.Fields{
            "customField": "SomeMessageCustomField",
	   },
   )
}
```

### Console log and output file

```go
package main
import (
    "time"

    "logging"
)

func main() {
   l := logging.InitLog("SomePackage", "tmp", "simple")
   ls := l.Log("SomeNameFunction")
   ls.SubLog(
       "info",
       "some message",
       time.Now().Format("2006-01-02T15:04:05"),
   ).Info("some info message")
   
   ls.SubLog(
       "warning",
       "some message",
       time.Now().Format("2006-01-02T15:04:05"),
   ).Warning("some warning message")
   
    ls.SubLog(
       "error",
       "some message",
       time.Now().Format("2006-01-02T15:04:05"),
    ).Error("some error message")
}
```

```bash
2020-06-02T17:28:15 [package] SomePackage [func] SomeNameFunction [INFO] some info message
2020-06-02T17:28:15 [package] SomePackage [func] SomeNameFunction [WARNING] some warning message
2020-06-02T17:28:15 [package] SomePackage [func] SomeNameFunction [ERROR] error message
```