package enum

type _levelLog struct {
	Error string
	Info  string
	Warn  string
	Debug string
}

//LoggerLevel enumeration
var LoggerLevel = &_levelLog{
	Error: "ERROR",
	Info:  "INFO",
	Warn:  "WARNING",
	Debug: "DEBUG",
}
