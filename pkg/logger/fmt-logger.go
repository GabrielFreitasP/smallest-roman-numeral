package logger

import (
	"fmt"
)

// Logger
type fmtLogger struct {
}

// App Logger constructor
func NewFmtLogger() Logger {
	return &fmtLogger{}
}

func (l *fmtLogger) InitLogger() {
}

// Logger methods

func (l *fmtLogger) Debug(args ...interface{}) {
	fmt.Println(args...)
}

func (l *fmtLogger) Debugf(template string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(template, args...))
}

func (l *fmtLogger) Info(args ...interface{}) {
	fmt.Println(args...)
}

func (l *fmtLogger) Infof(template string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(template, args...))
}

func (l *fmtLogger) Warn(args ...interface{}) {
	fmt.Println(args...)
}

func (l *fmtLogger) Warnf(template string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(template, args...))
}

func (l *fmtLogger) Error(args ...interface{}) {
	fmt.Println(args...)
}

func (l *fmtLogger) Errorf(template string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(template, args...))
}

func (l *fmtLogger) DPanic(args ...interface{}) {
	fmt.Println(args...)
}

func (l *fmtLogger) DPanicf(template string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(template, args...))
}

func (l *fmtLogger) Panic(args ...interface{}) {
	fmt.Println(args...)
}

func (l *fmtLogger) Panicf(template string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(template, args...))
}

func (l *fmtLogger) Fatal(args ...interface{}) {
	fmt.Println(args...)
}

func (l *fmtLogger) Fatalf(template string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(template, args...))
}
