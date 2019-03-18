package stopwatch

import (
	"time"
)

type Stopwatch struct {
	start    time.Time
	duration time.Duration
	running  bool
}

// NewStopwatch creates a new stopwatch object
func NewStopwatch() *Stopwatch {
	s := &Stopwatch{time.Time{}, 0, false}
	return s
}

// Duration returns the elapsed time for the stopwatch only counting
// time periods when the stopwatch is running.
func (s *Stopwatch) Duration() time.Duration {
	d := s.duration

	if true == s.running {
		now := DepGetTime()
		d += now.Sub(s.start)
	}

	return d
}

// Stop stops the stopwatch based on the current wall-clock time.
func (s *Stopwatch) Stop() time.Duration {
	s.duration = s.Duration()
	s.running = false

	return s.duration
}

// Start starts the stopwatch based on the current wall-clock time.
func (s *Stopwatch) Start() time.Duration {
	if false == s.running {
		s.running = true
		s.start = DepGetTime()
	}

	return s.duration
}

// Reset clears the stopwatch.
func (s *Stopwatch) Reset() time.Duration {
	s.running = false
	s.duration = 0

	return s.duration
}

// Milliseconds returns the elapsed duration in milliseconds.
func (s *Stopwatch) Milliseconds() int64 {
	return int64(s.Duration() * time.Nanosecond / time.Millisecond)
}

// Seconds returns the elapsed duration in seconds.
func (s *Stopwatch) Seconds() int64 {
	return int64(s.Duration() * time.Nanosecond / time.Second)
}

// Minutes returns the elapsed duration in minutes.
func (s *Stopwatch) Minutes() int64 {
	return int64(s.Duration() * time.Nanosecond / time.Minute)
}

// Hours returns the elapsed duration in hours.
func (s *Stopwatch) Hours() int64 {
	return int64(s.Duration() * time.Nanosecond / time.Hour)
}

// Days returns the elapsed duration in days.
func (s *Stopwatch) Days() int64 {
	return int64(s.Duration() * time.Nanosecond / (24 * time.Hour))
}
