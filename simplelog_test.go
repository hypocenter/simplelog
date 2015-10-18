package simplelog

import (
	"bytes"
	"testing"
)

func TestDefaultStdout(t *testing.T) {
	Trace("Some text printed by Trace()")
	Debug("Some text printed by Debug()")
	Info("Some text printed by Info()")
	Warning("Some text printed by Warning()")
	Error("Some text printed by Error()")
	Flush()
}

func TestDefaultReturnValue(t *testing.T) {
	err := Warning("Some text printed by Warning()")
	if err == nil || err.Error() != "[WARNING] Some text printed by Warning()" {
		t.Error("Warning() return wrong value.")
	}

	err = Error("Some text printed by Error()")
	if err == nil || err.Error() != "[ERROR] Some text printed by Error()" {
		t.Error("Error() return wrong value.")
	}
	Flush()
}

func TestCustomLog(t *testing.T) {
	buf := []byte{}
	writer := bytes.NewBuffer(buf)
	logger := New()
	logger.AddWriter(writer, L_DEBUG, W_BOTH)

}

func TestCustomLogWay(t *testing.T) {

}

func TestCustomLogPrefix(t *testing.T) {

}
