package utils

import (
	"time"

	"github.com/rs/zerolog/log"
)

type Timer struct {
	name        string
	totalTimer  time.Time
	sectionTime time.Time
}

func StartTimer(name string) Timer {
	log.Debug().Str("name", name).Msg("Starting Timer")
	return Timer{
		name:        name,
		totalTimer:  time.Now(),
		sectionTime: time.Now(),
	}
}

func (t *Timer) LogTime(segment string) {
	duration := time.Since(t.sectionTime)
	log.Debug().
		Str("name", t.name).
		Str("segment", segment).
		Str("duration", duration.String()).
		Msg("Timing")
	t.sectionTime = time.Now()
}

func (t *Timer) LogTotalTime() {
	duration := time.Since(t.totalTimer)
	log.Debug().
		Str("name", t.name).
		Str("segement", "Total Elapsed Time").
		Str("duration", duration.String()).
		Msg("Timing")
}
