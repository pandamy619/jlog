package logging

type Logs struct {
	location string
	Pkg string `json:"package"`
	Funs map[string]interface{} `json:"functions"`
	log []interface{}
	name string
}

type Fields map[string]interface{}

func InitLog(pkg string, location string) *Logs{
	m := make(map[string]interface{})
	return &Logs{location: location, Pkg: pkg, Funs: m}
}

func (l *Logs) Log(name string) *Logs{
	l.name = name
	return &Logs{location: l.location, Pkg: l.Pkg, Funs: l.Funs,  name: l.name, log: l.log}
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
	l.filedToConsole(fields)
	return &Logs{location: l.location, Pkg: l.Pkg, Funs: l.Funs,  name: l.name, log: l.log}
}

func (l *Logs) Report() []byte{
	obj := structJson(l)
	return obj
}
