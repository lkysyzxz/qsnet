package logger

import "log"

func Info(v string){
	log.Println("[INFO]:",v)
}

func Fatal(v string){
	log.Println("[ERROR]:",v)
}

type Logger interface{
	Info(v ...interface{})
	Fatal(v ...interface{})
}