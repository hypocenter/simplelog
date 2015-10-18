package simplelog

import (
	"fmt"
)

type log struct {
	// 要输出的内容
	str string
	// 用于替换的参数
	args []interface{}
	// 输出等级，L_TRACE ~ L_CRITICAL
	level int
}

func (l *log) output() string {
	str := l.str
	switch l.level {
	case L_TRACE:
		str = "[TRACE] " + str
	case L_DEBUG:
		str = "[DEBUG] " + str
	case L_INFO:
		str = "[INFO] " + str
	case L_WARNING:
		str = "[WARNING] " + str
	case L_ERROR:
		str = "[ERROR] " + str
	case L_CRITICAL:
		str = "[CRITICAL] " + str
	}

	if len(l.args) > 0 {
		str = fmt.Sprintf(str, l.args...) + "\n"
	}

	return str
}
