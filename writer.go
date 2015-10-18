package simplelog

import (
	slog "log"
)

type writer struct {
	lg *slog.Logger
	// 输出等级，L_TRACE ~ L_CRITICAL
	level int
	// 输出方式，W_INFO|W_ERROR
	way int
}

func (w *writer) write(l *log) {
	if (w.way&W_INFO == 0) && (l.level == L_TRACE || l.level == L_DEBUG || l.level == L_INFO) {
		return
	}

	if (w.way&W_ERROR == 0) && (l.level == L_WARNING || l.level == L_ERROR) {
		return
	}

	if l.level >= w.level {
		switch l.level {
		case L_TRACE, L_DEBUG, L_INFO, L_WARNING, L_ERROR:
			if len(l.args) > 0 {
				w.lg.Printf(l.str, l.args)
			} else {
				w.lg.Printf(l.str)
			}
		case L_CRITICAL:
			if len(l.args) > 0 {
				w.lg.Fatalf(l.str, l.args)
			} else {
				w.lg.Fatalf(l.str)
			}
		default:
			return
		}
	}
}
