# logger

## Installation
```bash
go get github.com/pandamy619/logger
```

## Usage

### Init struct

```go
NewLog(time string, out io.Writer)
```

### Print to console
```go
l := NewLog(time string, out io.Writer)

l.WithFields(Fields{
    "package": "SomePackage",
    "function": "SomeFunction",
})
```

#### Info
```go
l.Info("message info")

2020-05-23T11:08:00 [INFO] package: SomePackage function: SomeFunction message info
```

#### Warning
```go
l.Warning("message warning")

2020-05-23T11:08:00 [Warning] package: SomePackage function: SomeFunction message warning
```

#### Error
```go
l.Error("message error")

2020-05-23T11:08:00 [Error] package: SomePackage function: SomeFunction message error
```