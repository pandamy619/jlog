package jlog

type Fields map[string]interface{}

type Jlogs struct {
	location string
	Pkg string `json:"package"`
	Funs map[string]interface{} `json:"functions"`
	log []interface{}
	name string
	fields Fields
	consoleLog string
}

func Init(pkg string, location string, consoleLog string) *Jlogs {
	m := make(map[string]interface{})
	return &Jlogs{
		location: location,
		Pkg: pkg,
		Funs: m,
		consoleLog: consoleLog,
	}
}

func (l *Jlogs) Log(name string) *Jlogs {
	l.name = name
	return &Jlogs{
		location: l.location,
		Pkg: l.Pkg,
		Funs: l.Funs,
		name: l.name,
		log: l.log,
		fields: l.fields,
		consoleLog: l.consoleLog,
	}
}

func (l *Jlogs) SubLogWithFields(status string, message string, time string, fields Fields) *Jlogs {
	f := Fields{
		"status": status,
		"message": message,
		"time": time,
		"meta": fields,
	}
	l.fields = f
	return l.updateArr(f)
}

func (l *Jlogs) SubLog(status string, message string, time string) *Jlogs {
	f := Fields{
		"status": status,
		"message": message,
		"time": time,
	}
	l.fields = f
	return l.updateArr(f)
}

func (l *Jlogs) updateArr(fields Fields) *Jlogs {
	l.log = append(l.log, fields)
	l.Funs[l.name] = l.log
	return &Jlogs{
		location: l.location,
		Pkg: l.Pkg,
		Funs: l.Funs,
		name: l.name,
		log: l.log,
		fields: l.fields,
		consoleLog: l.consoleLog,
	}
}

func (l *Jlogs) Report() []byte{
	obj := structJson(l)
	return obj
}
