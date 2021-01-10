package jlog

type Jlogs struct {
	location    string
	Pkg         string                 `json:"package"`
	Funs        map[string]interface{} `json:"functions"`
	log         []interface{}
	name        string
	fields      Fields
	consoleType string
}

func Init(pkg string, format string, location string) *Jlogs {

	fun := make(map[string]interface{})
	return &Jlogs{
		location:    location,
		Pkg:         pkg,
		Funs:        fun,
		consoleType: format,
	}
}
