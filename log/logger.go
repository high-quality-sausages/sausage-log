package log

import (
	"fmt"
)

type Log struct {
	Connector
	Config
}

type Connector interface {
	Debug(fmtString string, args ...interface{})
	Info(fmtString string, args ...interface{})
	Warn(fmtString string, args ...interface{})
	Err(fmtString string, args ...interface{})
}

func Debug(fmtString string, args ...interface{}) {
	fmt.Print(fmtString)
}

func Info(fmtString string, args ...interface{}) {
	fmt.Println(fmtString)
}

func Warn(fmtString string, args ...interface{}) {}

func Err(fmtString string, args ...interface{}) {}

func (l Log) Info(fmtString string, args ...interface{}) {}

func (l Log) Debug(fmtString string, args ...interface{}) {}

func (l Log) Warn(fmtString string, args ...interface{}) {}

func (l Log) Err(fmtString string, args ...interface{}) {}
