package stopwatch

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStopwatchNew(t *testing.T) {
	sw := NewStopwatch()

	assert.NotNil(t, sw)
	assert.Equal(t, time.Duration(0), sw.duration)
	assert.Equal(t, false, sw.running)
	assert.Equal(t, true, sw.start.IsZero())
}

func TestStopwatch(t *testing.T) {
	Stubs()

	var result time.Duration

	increase := [5]time.Duration{
		1 * time.Minute,
		1 * time.Minute,
		1 * time.Minute,
		1 * time.Minute,
		1 * time.Minute,
	}

	expected := [5]time.Duration{
		(1 * time.Minute) * time.Nanosecond,
		(2 * time.Minute) * time.Nanosecond,
		(2 * time.Minute) * time.Nanosecond,
		(3 * time.Minute) * time.Nanosecond,
		(4 * time.Minute) * time.Nanosecond,
	}

	stop := [5]bool{
		false,
		false,
		true,
		false,
		false,
	}

	sw := NewStopwatch()
	sw.Start()

	for i := 0; i < len(increase); i++ {

		if true == stop[i] {
			sw.Stop()
		} else {
			sw.Start()
		}

		_stubClockAdvance(increase[i])

		result = sw.Duration()

		assert.Equal(t, expected[i], result)
	}

	StubsRestore()
}

func TestStopwatch_Reset(t *testing.T) {
	Stubs()

	var result time.Duration

	increase := [5]time.Duration{
		1 * time.Millisecond,
		1 * time.Second,
		1 * time.Minute,
		1 * time.Hour,
		24 * time.Hour,
	}

	expected := [5]time.Duration{
		(1 * time.Millisecond) * time.Nanosecond,
		(1 * time.Second) * time.Nanosecond,
		(1 * time.Minute) * time.Nanosecond,
		(1 * time.Hour) * time.Nanosecond,
		(24 * time.Hour) * time.Nanosecond,
	}

	sw := NewStopwatch()
	sw.Start()

	for i := 0; i < len(increase); i++ {

		sw.Reset()
		sw.Start()

		_stubClockAdvance(increase[i])

		result = sw.Duration()

		assert.Equal(t, expected[i], result)

		ms := int64((expected[i] / time.Millisecond))
		assert.Equal(t, ms, sw.Milliseconds())

		secs := int64((expected[i] / time.Second))
		assert.Equal(t, secs, sw.Seconds())

		mins := int64((expected[i] / time.Minute))
		assert.Equal(t, mins, sw.Minutes())

		hours := int64((expected[i] / time.Hour))
		assert.Equal(t, hours, sw.Hours())

		days := int64((expected[i] / (24 * time.Hour)))
		assert.Equal(t, days, sw.Days())
	}

	StubsRestore()
}

func TestStopwatch_Real(t *testing.T) {
	sw := NewStopwatch()
	sw.Start()

	time.Sleep(1 * time.Millisecond)

	result := int64(sw.Duration())
	assert.NotEqual(t, int64(0), result)
}

func ExampleStopwatch() {
	sw := NewStopwatch()
	sw.Start()

	time.Sleep(250 * time.Millisecond)

	fmt.Printf("Stopwatch Duration: %v\n", sw.Duration())

	sw.Stop()

	time.Sleep(250 * time.Millisecond)
	fmt.Printf("Stopwatch Duration: %v\n", sw.Duration())

	sw.Start()

	time.Sleep(250 * time.Millisecond)
	fmt.Printf("Stopwatch Duration: %v\n", sw.Duration())

	sw.Reset()
	sw.Start()

	time.Sleep(250 * time.Millisecond)
	fmt.Printf("Stopwatch Duration: %v\n", sw.Duration())
}
