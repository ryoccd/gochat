package logger

import (
	"fmt"
	"log"
	"os"
)

var Mlog *log.Logger

func init() {
	file, err := os.OpenFile("application.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Error("Failed to open log file", err)
	}

	Mlog = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Puts(a ...interface{}) {
	fmt.Println(a...)
}

func Info(a ...interface{}) {
	Mlog.SetPrefix("[INFO] ")
	Mlog.Println(a...)
}

func Error(a ...interface{}) {
	Mlog.SetPrefix("[Error] ")
	Mlog.Println(a...)
}

func Warn(a ...interface{}) {
	Mlog.SetPrefix("[Warn] ")
	Mlog.Println(a...)
}
