package logger

type Logger interface {
	Trace(string, string, interface{})
	Debug(string, string, interface{})
	Info(string, string, interface{})
	Warn(string, string, interface{})
	Error(string, string, interface{})
	Fatal(string, string, interface{})
	Panic(string, string, interface{})
}
