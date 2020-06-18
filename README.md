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

example
```bash
package main

import (
	"jlog"
)

func main() {
    jlog.Init("SomePackage", "tmp", "simple")
}
```

### func (*Jlogs) Main

```go
func (l *Jlogs) Main(name string) *Jlogs
```
Usually called at the beginning of a function.
* name: Function name.

example
```bash
package main

import (
	"jlog"
)

func main() {
    j := jlog.Init("SomePackage", "tmp", "simple")
    j.Main("main")
}
```

### func (*Jlogs) Sub

```go
func (l *Jlogs) Sub(status string, message string, time string) *Jlogs
```
* status: Status log (info/warning/error/customStatus).
* message: Message containing information about the log.
* time: Date and time.

```bash
package main

import (
	"jlog"
)

func main() {
    j := jlog.Init("SomePackage", "tmp", "simple")
    jMain := j.Main("main")
    jMain.Sub(
        "info",
        "message fun1",
        time.Now().Format("2006-01-02T15:04:05"),
    ).Info("Info message")
}
```

### func (*Jlogs) SubWithFields
```go
func (l *Jlogs) SubWithFields(status string, message string, time string, field Fields) *Jlogs
```
* status: Status log (info/warning/error/customStatus).
* message: Message containing information about the log.
* time: Date and time.
* fields: Fields.

example
```bash
package main

import (
	"jlog"
)

func main() {
    j := jlog.Init("SomePackage", "tmp", "simple")
    jMain := j.Main("main")
    lMain.SubWithFields(
        "warning",
        "message subfun1",
        time.Now().Format("2006-01-02T15:04:05"),
        jlog.Fields{
            "fun": "subfun1",
    }).Warning("Warning message")
}
```

#### Console log
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