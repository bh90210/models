package machine

import (
	"time"
)

func EqualDuration(x int, dur time.Duration) []time.Duration {
	var durations []time.Duration
	for range x {
		durations = append(durations, dur)
	}

	return durations
}
