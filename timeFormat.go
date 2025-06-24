package main

import (
	"github.com/mergestat/timediff"
	"time"
)

func toReadableTime(t *time.Time) string {
	return timediff.TimeDiff(*t)
}
