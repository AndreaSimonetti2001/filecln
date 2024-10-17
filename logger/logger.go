package logger

import "log"

func Info(format string, v ...any) {
	log.SetPrefix("INFO: ")
	log.Printf(format+"\n", v...)
}

func Warn(format string, v ...any) {
	log.SetPrefix("WARN: ")
	log.Printf(format+"\n", v...)
}

func Error(format string, v ...any) {
	log.SetPrefix("ERROR: ")
	log.Fatalf(format+"\n", v...)
}
