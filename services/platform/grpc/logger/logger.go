package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger() (*Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &Logger{logger: logger}, nil
}

func (l *Logger) Info(msg string, fields ...interface{}) {
	l.logger.Sugar().Infow(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...interface{}) {
	l.logger.Sugar().Errorw(msg, fields...)
}
