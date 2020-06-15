package lzap

//IService interface
type IService interface {
	SetLogLevel(level string)
	Printf(format string, v...interface{})
	Print(v...interface{})
	Info(msg string, tags ...string)
	Warn(msg string, tags ...string)
	Debug(msg string, tags...string)
	Error(msg string, err error, tags...string)
}
