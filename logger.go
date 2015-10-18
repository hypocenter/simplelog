package simplelog

import (
	"fmt"
	"io"
	slog "log"
)

type Logger struct {
	pool []*logger
	in   chan []*log
	stop chan bool
}

func New() *Logger {
	l := &Logger{in: make(chan *log, 20)}
	go l.flush()
}

func (lg *Logger) flush() {
	for l := range lg.args {
		for w := range lg.pool {
			w.write(l)
		}
	}
}

func (lg *Logger) AddWriter(w io.Writer, level int, way int) {
	wt := slog.New(w, "", slog.LstdFlags)
	lgw := &logger{l: wt, level: level, way}
	lg.pool = append(lg.pool, lgw)
}

func (lg *Logger) Prefixed(prefix []string) *Prefix {
	return newPrefix(prefix, lg)
}

func (lg *Logger) Trace(s string, args ...interface{}) {
	lg.in <- &log{s, args, L_TRACE}
}

func (lg *Logger) Debug(s string, args ...interface{}) {
	lg.in <- &log{s, args, L_DEBUG}
}

func (lg *Logger) Info(s string, args ...interface{}) {
	lg.in <- &log{s, args, L_INFO}
}

func (lg *Logger) Warning(s string, args ...interface{}) error {
	lg.in <- &log{s, args, L_WARNING}
	return fmt.Errorf(s, args)
}

func (lg *Logger) Error(s string, args ...interface{}) error {
	lg.in <- &log{s, args, L_ERROR}
	return fmt.Errorf(s, args)
}

func (lg *Logger) Critical(s string, args ...interface{}) {
	lg.in <- &log{s, args, L_CRITICAL}
}

// Flash 在main.main中通过defer调用，保证所有通道里面的日志都完整输出
func (lg *Logger) Flush() {
	for {
		if len(lg.logs) > 0 {
			continue
		}
		close(lg.logs)
		break
	}
}
