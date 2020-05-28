// package log provides implementation-agnostic easy-mockable interface
// "Logger" for a logger

package log

import (
	"os"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ Logger = logrus.FieldLogger(nil)
var _ Logger = &zap.SugaredLogger{}

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

type AsyncLogger interface {
	Logger

	Sync() error
}

func With(logger Logger, field string, value interface{}) Logger {
	switch logger := logger.(type) {
	case *zap.SugaredLogger:
		return logger.With(field, value)
	case *logrus.Entry:
		return logger.WithField(field, value)
	case *logrus.Logger:
		return logger.WithField(field, value)
	}
	return logger
}

func NewNop() Logger {
	return zap.NewNop().Sugar()
}

const (
	DebugLevel  = zap.DebugLevel
	InfoLevel   = zap.InfoLevel
	WarnLevel   = zap.WarnLevel
	ErrorLevel  = zap.ErrorLevel
	DPanicLevel = zap.DPanicLevel
	PanicLevel  = zap.PanicLevel
	FatalLevel  = zap.FatalLevel
)

type Level = zapcore.Level

func New(loggingLevelFlag Level) AsyncLogger {
	return zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.Lock(os.Stdout),
		zap.NewAtomicLevelAt(loggingLevelFlag),
	)).Sugar()
}

func LevelFlag(name string, defaultLevel Level, usage string) *Level {
	return zap.LevelFlag(name, defaultLevel, usage)
}
