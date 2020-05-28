package log

var _ Logger = MockLogger{}

type MockLogger struct {
	FuncDebugf func(format string, args ...interface{})
	FuncInfof  func(format string, args ...interface{})
	FuncWarnf  func(format string, args ...interface{})
	FuncErrorf func(format string, args ...interface{})
	FuncPanicf func(format string, args ...interface{})
	FuncFatalf func(format string, args ...interface{})
	FuncSync   func() error
}

func (logger MockLogger) Debugf(format string, args ...interface{}) {
	logger.FuncDebugf(format, args...)
}

func (logger MockLogger) Infof(format string, args ...interface{}) {
	logger.FuncInfof(format, args...)
}

func (logger MockLogger) Warnf(format string, args ...interface{}) {
	logger.FuncWarnf(format, args...)
}

func (logger MockLogger) Errorf(format string, args ...interface{}) {
	logger.FuncErrorf(format, args...)
}

func (logger MockLogger) Panicf(format string, args ...interface{}) {
	logger.FuncPanicf(format, args...)
}

func (logger MockLogger) Fatalf(format string, args ...interface{}) {
	logger.FuncFatalf(format, args...)
}

func (logger MockLogger) Sync() error {
	return logger.FuncSync()
}
