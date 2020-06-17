# jlog

## Installation
```bash
go get github.com/pandamy619/logger
```

## Usage

### func Init

```go
func Init(pkg string, location string, consoleLog string) *Jlogs
```
Create new Logger.
* pkg: Name package.
* location: Folder where logs will be saved.
* consoleLog: Method output to console, simple or json.

### func (*Jlogs) Main

```go
func (l *Jlogs) Main(name string) *Jlogs
```
Usually called at the beginning of a function.
* name: Function name.

### func (*Jlogs) Sub

```go
func (l *Jlogs) Sub(status string, message string, time string) *Jlogs
```
* status: Status log (info/warning/error/customStatus).
* message: Message containing information about the log.
* time: Date and time.

### func (*Jlogs) SubWithFields
```go
func (l *Jlogs) SubLogWithFields(status string, message string, time string, field Fields) *Jlogs
```
* status: Status log (info/warning/error/customStatus).
* message: Message containing information about the log.
* time: Date and time.
* fields: Fields.

## Example

```go
package main
import (
    "time"

    "jlog"
)

func fun1(l *Jlogs) {
   l.Main("fun1").Sub(
   	   "info",
       "some message",
       time.Now().Format("2006-01-02T15:04:05"),
   ).Info("some info message")
}

func fun2(l *Jlogs) {
   ls := l.Main("fun2")
   ls.Sub(
       "warning",
       "some message",
       time.Now().Format("2006-01-02T15:04:05"),
   ).Warning("some warning message")
}

func fun3(l *Jlogs) {
    ls := l.Main("fun3")
    ls.Sub(
       "error",
       "some message",
       time.Now().Format("2006-01-02T15:04:05"),
    ).Error("some error message")
}

func main() {
   l := jlog.Init("SomePackage", "tmp", "simple")
   fun1(l)
   fun2(l)
   fun3(l)
}
```

simple log
```bash
[INFO] package:some_package func:fun1 detail:Info message
[WARN] package:some_package func:fun1 detail:warning message
[ERROR] package:some_package func:fun1 detail:error message
```

json log

```bash
-------------------------
[INFO] Info message
package:some_package
func:fun1
{
    "message": "message fun1",
    "status": "info",
    "time": "2020-06-17T11:21:59"
}
-------------------------
-------------------------
[WARN] warning message
package:some_package
func:fun1
{
    "message": "message fun1",
    "status": "warning",
    "time": "2020-06-17T11:21:59"
}
-------------------------
-------------------------
[ERROR] error message
package:some_package
func:fun1
{
    "message": "message fun1",
    "status": "error",
    "time": "2020-06-17T11:21:59"
}
-------------------------
```