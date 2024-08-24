package utils

import (
	"log/slog"
	"time"
)

type Timer struct {
	totalTimer  time.Time
	sectionTime time.Time
}

func StartTimer(message string) Timer {
	slog.Info("Starting to time: " + message)
	return Timer{
		totalTimer:  time.Now(),
		sectionTime: time.Now(),
	}
}

func (t *Timer) LogTime(message string) {
	slog.Info(message, "Segement Elapsed Time", time.Duration(time.Since(t.sectionTime)).String())
	t.sectionTime = time.Now()
}

func (t *Timer) LogTotalTime(message string) {
	slog.Info(message, "Total Elapsed Time", time.Duration(time.Since(t.totalTimer)).String())
}
