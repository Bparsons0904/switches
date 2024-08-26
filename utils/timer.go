package utils

import (
	"fmt"
	"log/slog"
	"time"
)

type Timer struct {
	name        string
	totalTimer  time.Time
	sectionTime time.Time
}

func StartTimer(name string) Timer {
	slog.Info("Starting to time: " + name)
	return Timer{
		name:        name,
		totalTimer:  time.Now(),
		sectionTime: time.Now(),
	}
}

func (t *Timer) LogTime(segment string) {
	title := fmt.Sprintf("Segment %s Elapsed Time", segment)
	slog.Info(t.name, title, time.Duration(time.Since(t.sectionTime)).String())
	t.sectionTime = time.Now()
}

func (t *Timer) LogTotalTime() {
	slog.Info(t.name, "Total Elapsed Time", time.Duration(time.Since(t.totalTimer)).String())
}
