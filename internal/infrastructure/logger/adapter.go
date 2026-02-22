package logger

import (
	"github.com/BT2701/backend-fishing-gameplay/internal/infrastructure/contract"
	"go.uber.org/zap"
)

// ZapLoggerAdapter wraps zap.Logger to implement contract.Logger interface
type ZapLoggerAdapter struct {
	logger *zap.Logger
}

// NewZapLoggerAdapter creates a new logger adapter
func NewZapLoggerAdapter(zapLogger *zap.Logger) contract.Logger {
	return &ZapLoggerAdapter{
		logger: zapLogger,
	}
}

func (z *ZapLoggerAdapter) Debug(msg string, fields ...interface{}) {
	z.logger.Sugar().Debugw(msg, fields...)
}

func (z *ZapLoggerAdapter) Info(msg string, fields ...interface{}) {
	z.logger.Sugar().Infow(msg, fields...)
}

func (z *ZapLoggerAdapter) Warn(msg string, fields ...interface{}) {
	z.logger.Sugar().Warnw(msg, fields...)
}

func (z *ZapLoggerAdapter) Error(msg string, fields ...interface{}) {
	z.logger.Sugar().Errorw(msg, fields...)
}

func (z *ZapLoggerAdapter) Fatal(msg string, fields ...interface{}) {
	z.logger.Sugar().Fatalw(msg, fields...)
}

func (z *ZapLoggerAdapter) Sync() error {
	return z.logger.Sync()
}
