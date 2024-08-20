package utils

import (
	"log"
	"time"
)

type Timer struct {
	totalTimer  time.Time
	sectionTime time.Time
}

func StartTimer(message string) Timer {
	log.Println("Starting to time: ", message)
	return Timer{
		totalTimer:  time.Now(),
		sectionTime: time.Now(),
	}
}

func (t *Timer) LogTime(message string) {
	log.Println(message, time.Duration(time.Since(t.sectionTime)).String())
	t.sectionTime = time.Now()
}

func (t *Timer) LogTotalTime(message string) {
	log.Println(message, time.Duration(time.Since(t.totalTimer)).String())
}

var (
	totalTimer  time.Time
	sectionTime time.Time
)

func StartTimer2() {
	sectionTime = time.Now()
	totalTimer = sectionTime
}

func LogTime2(message string) {
	log.Println(message, time.Duration(time.Since(sectionTime)).String())
	sectionTime = time.Now()
}

func LogTotalTime2(message string) {
	log.Println(message, time.Duration(time.Since(totalTimer)).String())
}
