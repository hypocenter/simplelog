package simplelog

type log struct {
	// 要输出的内容
	str string
	// 用于替换的参数
	args []interface{}
	// 输出等级，L_TRACE ~ L_CRITICAL
	level int
}
