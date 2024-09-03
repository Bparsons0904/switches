package logutil

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/gorm/logger"
)

type GormLogger struct {
	ZeroLogger zerolog.Logger
	LogLevel   logger.LogLevel
}

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.ZeroLogger.Info().Msgf(msg, data...)
	}
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.ZeroLogger.Warn().Msgf(msg, data...)
	}
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.ZeroLogger.Error().Msgf(msg, data...)
	}
}

func (l *GormLogger) Trace(
	ctx context.Context,
	begin time.Time,
	fc func() (string, int64),
	err error,
) {
	if l.LogLevel > logger.Silent {
		elapsed := time.Since(begin)
		sql, rows := fc()

		logEvent := l.ZeroLogger.Debug()
		if err != nil {
			logEvent = l.ZeroLogger.Error().Err(err)
		}

		logEvent.
			Str("sql", sql).
			Int64("rows", rows).
			Dur("elapsed", elapsed).
			Msg("GORM Query")
	}
}

// NewGormLogger creates a new GORM logger that uses zerolog
func NewGormLogger(level logger.LogLevel) *GormLogger {
	return &GormLogger{
		ZeroLogger: GetLogger(),
		LogLevel:   level,
	}
}
