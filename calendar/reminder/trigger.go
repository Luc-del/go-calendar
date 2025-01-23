package reminder

import (
	"fmt"
	"time"
)

const (
	Relative = "relative"
)

type Timing bool

const (
	Before Timing = false
	After  Timing = true
)

type RelativeTrigger struct {
	Timing
	time.Duration
}

func (t RelativeTrigger) render() string {
	totalSeconds := int(t.Duration.Seconds())

	// Calculate days, hours, minutes, and seconds
	days := totalSeconds / (24 * 3600)
	totalSeconds %= 24 * 3600
	hours := totalSeconds / 3600
	totalSeconds %= 3600
	minutes := totalSeconds / 60
	seconds := totalSeconds % 60

	// Build the ICS duration format string
	res := "P"
	if days > 0 {
		res += fmt.Sprintf("%dD", days)
	}
	if hours > 0 || minutes > 0 || seconds > 0 {
		res += "T"
		if hours > 0 {
			res += fmt.Sprintf("%dH", hours)
		}
		if minutes > 0 {
			res += fmt.Sprintf("%dM", minutes)
		}
		if seconds > 0 {
			res += fmt.Sprintf("%dS", seconds)
		}
	}

	// If no time components were added, default to PT0S to represent zero duration
	if res == "P" {
		res += "T0S"
	}

	// Handle timing
	if t.Timing == Before {
		res = "-" + res
	}
	return "TRIGGER:" + res
}

// TODO Absolute time reminder trigger. Ex: TRIGGER;VALUE=DATE-TIME:20250123T090000Z
