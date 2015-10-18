package simplelog

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"time"
)

func TestDefaultStdout(t *testing.T) {
	Trace("Some text printed by Trace()")
	Debug("Some text printed by Debug()")
	Info("Some text printed by Info()")
	err := Warning("Some text printed by Warning()")
	if err == nil || err.Error() != "[WARNING] Some text printed by Warning()" {
		t.Error("Warning() return wrong value.")
	}
	err = Error("Some text printed by Error()")
	if err == nil || err.Error() != "[ERROR] Some text printed by Error()" {
		t.Error("Error() return wrong value.")
	}

	root.restart()
}

func TestCustomLog(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.AddWriter(buf, L_TRACE, W_BOTH)

	tests := map[string]func(string, ...interface{}){
		"TRACE": logger.Trace,
		"DEBUG": logger.Debug,
		"INFO":  logger.Info,
	}

	for n, f := range tests {
		f("Some %s passed to "+n+"()", "text")
		logger.restart()
		now := time.Now()
		expect := now.Format("2006/01/02 15:04:05") + " [" + n + "] Some text passed to " + n + "()\n"
		res, err := buf.ReadBytes('\n')
		if err != nil && err != io.EOF {
			t.Error(err)
		}
		if string(res) != expect {
			t.Errorf(n+"() printed wrong log expect '%s' got '%s'\n", expect, res)
		}
	}

	tests2 := map[string]func(string, ...interface{}) error{
		"WARNING": logger.Warning,
		"ERROR":   logger.Error,
	}

	for n, f := range tests2 {
		f("Some %s passed to "+n+"()", "text")
		logger.restart()
		now := time.Now()
		expect := now.Format("2006/01/02 15:04:05") + " [" + n + "] Some text passed to " + n + "()\n"
		res, err := buf.ReadBytes('\n')
		if err != nil && err != io.EOF {
			t.Error(err)
		}
		if string(res) != expect {
			t.Errorf(n+"() printed wrong log expect '%s' got '%s'\n", expect, res)
		}
	}

}

func TestCustomLogLevel(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.AddWriter(buf, L_WARNING, W_BOTH)
	logger.Trace("Some text from Trace()")
	logger.Debug("Some text from DEBUG()")
	logger.Info("Some text from INFO()")
	logger.restart()

	res, err := buf.ReadBytes('\n')
	if err != io.EOF {
		t.Error("Log to wrong level")
	}

	logger.Warning("Some text from Warning()")
	logger.restart()
	res, err = buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		t.Error(err)
	}
	if len(res) < 0 {
		t.Error("Log to wrong level")
	}

	logger.Error("Some text from Error()")
	logger.restart()
	res, err = buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		t.Error(err)
	}
	if len(res) < 0 {
		t.Error("Log to wrong level")
	}
}

func TestCustomLogWayInfo(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.AddWriter(buf, L_TRACE, W_INFO)

	logger.Warning("Some text from Warning()")
	logger.restart()
	res, err := buf.ReadBytes('\n')
	if err != io.EOF {
		t.Error("Log to wrong way")
	}

	logger.Error("Some text from Error()")
	logger.restart()
	res, err = buf.ReadBytes('\n')
	if err != io.EOF {
		t.Error("Log to wrong way")
	}

	logger.Trace("Some text from Trace()")
	logger.restart()
	res, err = buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		t.Error(err)
	}
	if len(res) < 0 {
		t.Error("Log to wrong way")
	}

	logger.Info("Some text from Info()")
	logger.restart()
	res, err = buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		t.Error(err)
	}
	if len(res) < 0 {
		t.Error("Log to wrong way")
	}
}

func TestCustomLogWayError(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.AddWriter(buf, L_WARNING, W_ERROR)

	logger.Trace("Some text from Trace()")
	logger.restart()
	res, err := buf.ReadBytes('\n')
	if err != io.EOF {
		t.Error("Log to wrong way")
	}

	logger.Debug("Some text from Debug()")
	logger.restart()
	res, err = buf.ReadBytes('\n')
	if err != io.EOF {
		t.Error("Log to wrong way")
	}

	logger.Error("Some text from Error()")
	logger.restart()
	res, err = buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		t.Error(err)
	}
	if len(res) < 0 {
		t.Error("Log to wrong way")
	}

	logger.Warning("Some text from Warning()")
	logger.restart()
	res, err = buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		t.Error(err)
	}
	if len(res) < 0 {
		t.Error("Log to wrong way")
	}
}

func TestCustomLogPrefix(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.AddWriter(buf, L_TRACE, W_BOTH)
	l := logger.Prefixed([]string{"HYPO", "CINDY"})
	l.Info("Some %s passed to Info()", "text")
	logger.restart()
	now := time.Now()
	expect := now.Format("2006/01/02 15:04:05") + " [INFO] [HYPO|CINDY] Some text passed to Info()\n"
	res, err := buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		t.Error(err)
	}
	if string(res) != expect {
		t.Errorf("Logger wrote wrong prefix expect '%s' got '%s'", expect, res)
	}
}

func TestCustomLogPrefixOpreation(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.AddWriter(buf, L_TRACE, W_BOTH)
	l := logger.Prefixed([]string{"HYPO"})
	jp := func(s []string) string {
		return P_STAR + strings.Join(s, P_SEPE) + P_END
	}
	if l.prefix != jp([]string{"HYPO"}) {
		t.Errorf("Prefix generated wrong prefix expect '%s' got '%s'", jp([]string{"HYPO"}), l.prefix)
	}

	l.AppendPrefix("CINDY")
	if l.prefix != jp([]string{"HYPO", "CINDY"}) {
		t.Errorf("Prefix generated wrong prefix got '%s'", l.prefix)
	}

	l.PrependPrefix("HU")
	if l.prefix != jp([]string{"HU", "HYPO", "CINDY"}) {
		t.Errorf("Prefix generated wrong prefix got '%s'", l.prefix)
	}

	l.CleanPrefix()
	if l.prefix != jp([]string{""}) {
		t.Errorf("Prefix generated wrong prefix got '%s'", l.prefix)
	}

	l.PrependPrefix("HYPO")
	if l.prefix != jp([]string{"HYPO"}) {
		t.Errorf("Prefix generated wrong prefix got '%s'", l.prefix)
	}

	l.CleanPrefix()
	l.AppendPrefix("HYPO")
	if l.prefix != jp([]string{"HYPO"}) {
		t.Errorf("Prefix generated wrong prefix got '%s'", l.prefix)
	}
}

func TestMuiltLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.AddWriter(buf, L_TRACE, W_BOTH)
	logger.AddWriter(buf, L_TRACE, W_BOTH)
	logger.Info("Some text passed to Info()")
	logger.restart()

	now := time.Now()
	expect := now.Format("2006/01/02 15:04:05") + " [INFO] Some text passed to Info()\n"
	res, err := buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		t.Error(err)
	}
	if string(res) != expect {
		t.Errorf("Logger print wrong log expect '%s' got '%s'", expect, res)
	}

	res, err = buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		t.Error(err)
	}
	if string(res) != expect {
		t.Errorf("Logger print wrong log expect '%s' got '%s'", expect, res)
	}
}
