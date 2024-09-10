package logger

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(logger *zap.Logger) *ZapLogger {
	return &ZapLogger{
		logger: logger,
	}
}

func (l *ZapLogger) Debug(msg string, fields ...interface{}) {
	l.logger.Sugar().Debugw(msg, fields...)
}

func (l *ZapLogger) Info(msg string, fields ...interface{}) {
	l.logger.Sugar().Infow(msg, fields...)
}

func (l *ZapLogger) Warn(msg string, fields ...interface{}) {
	l.logger.Sugar().Warnw(msg, fields...)
}

func (l *ZapLogger) Error(msg string, fields ...interface{}) {
	l.logger.Sugar().Errorw(msg, fields...)
}

func (l *ZapLogger) Fatal(msg string, fields ...interface{}) {
	l.logger.Sugar().Fatalw(msg, fields...)
}
