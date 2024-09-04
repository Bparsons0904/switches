package logutil

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var globalLogger zerolog.Logger

func init() {
	env := os.Getenv("TIER")

	var output io.Writer
	if env == "" || env == "development" {
		output = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "2006/01/02 15:04:05",
			NoColor:    false,
			FormatTimestamp: func(i interface{}) string {
				switch t := i.(type) {
				case string:
					if parsed, err := time.Parse(time.RFC3339, t); err == nil {
						return parsed.Format("2006/01/02 15:04:05")
					}
					return t
				case time.Time:
					return t.Format("2006/01/02 15:04:05")
				default:
					return fmt.Sprint(i)
				}
			},
		}
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		output = os.Stdout
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	}

	globalLogger = zerolog.New(output).With().Timestamp().Logger()
	log.Logger = globalLogger
}

func GetLogger() zerolog.Logger {
	return globalLogger
}

type GormLogger struct {
	ZeroLogger    zerolog.Logger
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
}

func NewGormLogger(zlog zerolog.Logger) *GormLogger {
	logLevel := logger.Info
	if os.Getenv("TIER") != "development" {
		logLevel = logger.Warn
	}
	return &GormLogger{
		ZeroLogger:    zlog,
		LogLevel:      logLevel,
		SlowThreshold: 200 * time.Millisecond,
	}
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
	if l.LogLevel <= logger.Info {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error && !errors.Is(err, gorm.ErrRecordNotFound):
		sql, rows := fc()
		l.ZeroLogger.Error().
			Err(err).
			Str("sql", sql).
			Int64("rows", rows).
			Dur("elapsed", elapsed).
			Msg("trace")
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		l.ZeroLogger.Warn().
			Str("sql", sql).
			Int64("rows", rows).
			Dur("elapsed", elapsed).
			Msg("slow query")
	case l.LogLevel == logger.Info:
		sql, rows := fc()
		l.ZeroLogger.
			Trace().
			Str("sql", sql).
			Int64("rows", rows).
			Dur("elapsed", elapsed).
			Msg("trace")
	}
}
