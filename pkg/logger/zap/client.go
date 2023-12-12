package zap

import (
	"errors"
	"os"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type client struct {
	level string
	caller,
	stacktrace bool
	logger *zap.SugaredLogger
}

var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// NewLogger instance logger
func NewLogger(level string, caller, stacktrace bool) *client {
	return &client{level, caller, stacktrace, nil}
}

func (l *client) InitLogger() {
	logLevel := l.getLoggerLevel(l.level)

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// write syncers
	stderrSyncer := zapcore.Lock(os.Stderr)

	l.logger = zap.New(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg),
			stderrSyncer,
			zap.NewAtomicLevelAt(logLevel)),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.LevelOf(zap.ErrorLevel)),
		zap.AddCallerSkip(1)).
		Sugar()

	if err := l.logger.Sync(); err != nil && !errors.Is(err, syscall.ENOTTY) {
		l.logger.Error(err)
	}
}

func (l *client) getLoggerLevel(lv string) zapcore.Level {
	level, exist := loggerLevelMap[lv]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

func (l *client) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *client) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *client) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *client) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}
