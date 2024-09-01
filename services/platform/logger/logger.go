package logger

import (
	"github.com/popeskul/awesome-messanger/services/platform/logger/ports"
	"go.uber.org/zap"
)

type zapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(logger *zap.Logger) ports.Logger {
	return &zapLogger{
		logger: logger,
	}
}

func (l *zapLogger) Debug(msg string, fields ...interface{}) {
	l.logger.Sugar().Debugw(msg, fields...)
}

func (l *zapLogger) Info(msg string, fields ...interface{}) {
	l.logger.Sugar().Infow(msg, fields...)
}

func (l *zapLogger) Warn(msg string, fields ...interface{}) {
	l.logger.Sugar().Warnw(msg, fields...)
}

func (l *zapLogger) Error(msg string, fields ...interface{}) {
	l.logger.Sugar().Errorw(msg, fields...)
}

func (l *zapLogger) Fatal(msg string, fields ...interface{}) {
	l.logger.Sugar().Fatalw(msg, fields...)
}
