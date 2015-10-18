package simplelog

import (
	"os"
)

const (
	L_TRACE = iota
	L_DEBUG
	L_INFO
	L_WARNING
	L_ERROR
	L_CRITICAL
)

const (
	W_INFO = 1 << iota
	W_ERROR
	W_BOTH = W_INFO | W_ERROR
)

const (
	P_STAR = "["
	P_SEPE = "|"
	P_END  = "] "
)

var (
	root = New()
)

func init() {
	root.AddWriter(os.Stdout, L_TRACE)
}

func Trace(s string, args ...interface{}) {
	root.Trace(s, args...)
}

func Debug(s string, args ...interface{}) {
	root.Debug(s, args...)
}

func Info(s string, args ...interface{}) {
	root.Info(s, args...)
}

func Warning(s string, args ...interface{}) error {
	return root.Warning(s, args...)
}

func Error(s string, args ...interface{}) error {
	return root.Error(s, args...)
}

func Critical(s string, args ...interface{}) {
	root.Critical(s, args...)
}
