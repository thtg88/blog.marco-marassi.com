package timer

import "time"

type Timer interface {
	Now() time.Time
	Since(t time.Time) time.Duration
}

// RealTimer provides a Timer interface implementation that uses the actual `time` package.
// This is especially useful when we need to swap implementation to a fake timer during tests.
type RealTimer struct{}

func (rt *RealTimer) Now() time.Time {
	return time.Now()
}

func (rt *RealTimer) Since(t time.Time) time.Duration {
	return time.Since(t)
}
