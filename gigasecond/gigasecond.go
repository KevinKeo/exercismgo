// Package gigasecond which add a gigasecond to a time.Time
package gigasecond

import (
	"time"
)

// Constant declaration.
const testVersion = 4

// AddGigasecond return a time plus gigasecond
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(1e9) * time.Second)
}
