package simplelog

import (
	"fmt"
	"io"
	slog "log"
)

type Logger struct {
	pool []*writer
	in   chan *log
}

func New() *Logger {
	lg := &Logger{in: make(chan *log, 20)}
	go lg.flush()
	loggerContainer = append(loggerContainer, lg)
	return lg
}

func (lg *Logger) flush() {
	for l := range lg.in {
		for _, w := range lg.pool {
			w.write(l)
		}
	}
}

func (lg *Logger) AddWriter(w io.Writer, level int, way int) {
	slg := slog.New(w, "", slog.LstdFlags)
	lgw := &writer{lg: slg, level: level, way: way}
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
	if len(args) > 0 {
		return fmt.Errorf(s, args)
	} else {
		return fmt.Errorf(s)
	}
}

func (lg *Logger) Error(s string, args ...interface{}) error {
	lg.in <- &log{s, args, L_ERROR}
	if len(args) > 0 {
		return fmt.Errorf(s, args)
	} else {
		return fmt.Errorf(s)
	}
}

func (lg *Logger) Critical(s string, args ...interface{}) {
	lg.in <- &log{s, args, L_CRITICAL}
}
