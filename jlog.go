package jlog

func Init(pkg string, location string, consoleType string) *Jlogs {
	fun := make(map[string]interface{})
	return &Jlogs{
		location:    location,
		Pkg:         pkg,
		Funs:        fun,
		consoleType: consoleType,
	}
}

type Jlogs struct {
	location    string
	Pkg         string `json:"package"`
	Funs        map[string]interface{} `json:"functions"`
	log         []interface{}
	name        string
	fields      Fields
	consoleType string
}

func (l *Jlogs) Main(name string) *Jlogs {
	l.name = name
	return &Jlogs{
		location:    l.location,
		Pkg:         l.Pkg,
		Funs:        l.Funs,
		name:        l.name,
		log:         l.log,
		fields:      l.fields,
		consoleType: l.consoleType,
	}
}

func (l *Jlogs) Sub(status string, message string, time string) *Jlogs {
	f := Fields{
		"status": status,
		"message": message,
		"time": time,
	}
	l.fields = f
	return l.updateArr(f)
}

func (l *Jlogs) SubWithFields(status string, message string, time string, fields Fields) *Jlogs {
	f := Fields{
		"status": status,
		"message": message,
		"time": time,
		"meta": fields,
	}
	l.fields = f
	return l.updateArr(f)
}

func (l *Jlogs) updateArr(fields Fields) *Jlogs {
	l.log = append(l.log, fields)
	l.Funs[l.name] = l.log
	return &Jlogs{
		location:    l.location,
		Pkg:         l.Pkg,
		Funs:        l.Funs,
		name:        l.name,
		log:         l.log,
		fields:      l.fields,
		consoleType: l.consoleType,
	}
}

func (l *Jlogs) Report() []byte{
	return convertToJson(l)
}