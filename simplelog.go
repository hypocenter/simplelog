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
	root            = New()
	loggerContainer []*Logger
)

func init() {
	root.AddWriter(os.Stdout, L_TRACE, W_BOTH)
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

// Flash 在main.main中通过defer调用，保证所有通道里面的日志都完整输出
func Flush() {
	for {
		for _, lg := range loggerContainer {
			for {
				if len(lg.in) > 0 {
					continue
				}
				// close(lg.in)
				break
			}
		}
		break
	}
}
