package utils

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

type Timer struct {
	name        string
	totalTimer  time.Time
	sectionTime time.Time
}

func StartTimer(name string) Timer {
	log.Info().Str("name", name).Msg("Starting Timer")
	return Timer{
		name:        name,
		totalTimer:  time.Now(),
		sectionTime: time.Now(),
	}
}

func (t *Timer) LogTime(segment string) {
	title := fmt.Sprintf("Segment %s Elapsed Time", segment)
	log.Info().
		Str("name", t.name).
		Str("title", title).
		Str("time", time.Duration(time.Since(t.sectionTime)).String()).
		Msg("Timing")
	t.sectionTime = time.Now()
}

func (t *Timer) LogTotalTime() {
	log.Info().
		Str("name", t.name).
		Str("title", "Total Elapsed Time").
		Str("time", time.Duration(time.Since(t.totalTimer)).String()).
		Msg("Timing")
}
