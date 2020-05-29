package log

type Printfer struct {
	Func func(format string, args ...interface{})
}

func (p Printfer) Printf(format string, args ...interface{}) {
	p.Func(format, args...)
}
