package simplelog

import (
	"strings"
)

type Prefix struct {
	prefix string
	lg     *Logger
}

func newPrefix(prefix []string, lg *Logger) *Prefix {
	p := P_STAR + strings.Join(prefix, P_SEPE) + P_END
	return &Prefix{prefix: p, lg: lg}
}

func (p *Prefix) AppendPrefix(prefix string) {
	old := strings.Split(p.prefix[len(P_STAR):len(p.prefix)-len(P_END)], P_SEPE)
	old = cleanEmptyPrefixSlice(old)
	p.prefix = P_STAR + strings.Join(append(old, prefix), P_SEPE) + P_END
}

func (p *Prefix) PrependPrefix(prefix string) {
	old := strings.Split(p.prefix[len(P_STAR):len(p.prefix)-len(P_END)], P_SEPE)
	old = cleanEmptyPrefixSlice(old)
	p.prefix = P_STAR + strings.Join(append([]string{prefix}, old...), P_SEPE) + P_END
}

func (p *Prefix) CleanPrefix() {
	p.prefix = P_STAR + P_END
}

func (p *Prefix) Trace(s string, args ...interface{}) {
	p.lg.Trace(p.prefix+s, args...)
}

func (p *Prefix) Debug(s string, args ...interface{}) {
	p.lg.Debug(p.prefix+s, args...)
}

func (p *Prefix) Info(s string, args ...interface{}) {
	p.lg.Info(p.prefix+s, args...)
}

func (p *Prefix) Warning(s string, args ...interface{}) error {
	return p.lg.Warning(p.prefix+s, args...)
}

func (p *Prefix) Error(s string, args ...interface{}) error {
	return p.lg.Error(p.prefix+s, args...)
}

func (p *Prefix) Critical(s string, args ...interface{}) {
	p.lg.Trace(p.prefix+s, args...)
}

func cleanEmptyPrefixSlice(s []string) []string {
	r := []string{}
	for _, v := range s {
		if v != "" {
			r = append(r, v)
		}
	}
	return r
}
