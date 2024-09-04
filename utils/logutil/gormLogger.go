package logutil

//
// import (
// 	"context"
// 	"errors"
// 	"time"
//
// 	"github.com/rs/zerolog"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )
//
// type GormLogger struct {
// 	ZeroLogger    zerolog.Logger
// 	LogLevel      logger.LogLevel
// 	SlowThreshold time.Duration
// }
//
// func NewGormLogger(zlog zerolog.Logger) *GormLogger {
// 	return &GormLogger{
// 		ZeroLogger:    zlog,
// 		LogLevel:      logger.Info,
// 		SlowThreshold: 200 * time.Millisecond,
// 	}
// }
//
// func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
// 	newLogger := *l
// 	newLogger.LogLevel = level
// 	return &newLogger
// }
//
// func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
// 	if l.LogLevel >= logger.Info {
// 		l.ZeroLogger.Info().Msgf(msg, data...)
// 	}
// }
//
// func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
// 	if l.LogLevel >= logger.Warn {
// 		l.ZeroLogger.Warn().Msgf(msg, data...)
// 	}
// }
//
// func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
// 	if l.LogLevel >= logger.Error {
// 		l.ZeroLogger.Error().Msgf(msg, data...)
// 	}
// }
//
// func (l *GormLogger) Trace(
// 	ctx context.Context,
// 	begin time.Time,
// 	fc func() (string, int64),
// 	err error,
// ) {
// 	if l.LogLevel <= logger.Silent {
// 		return
// 	}
//
// 	elapsed := time.Since(begin)
// 	switch {
// 	case err != nil && l.LogLevel >= logger.Error && !errors.Is(err, gorm.ErrRecordNotFound):
// 		sql, rows := fc()
// 		l.ZeroLogger.Error().
// 			Err(err).
// 			Str("sql", sql).
// 			Int64("rows", rows).
// 			Dur("elapsed", elapsed).
// 			Msg("trace")
// 	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
// 		sql, rows := fc()
// 		l.ZeroLogger.Warn().
// 			Str("sql", sql).
// 			Int64("rows", rows).
// 			Dur("elapsed", elapsed).
// 			Msg("slow query")
// 	case l.LogLevel == logger.Info:
// 		_, rows := fc()
// 		l.ZeroLogger.Info().Int64("rows", rows).Dur("elapsed", elapsed).Msg("trace")
// 	}
// }
