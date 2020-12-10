package retry

import "time"

type Retry struct {
	MaxAttempts int
	Interval    time.Duration
}
