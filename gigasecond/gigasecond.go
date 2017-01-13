// Package gigasecond.
package gigasecond

import (
	"time"
)

// Constant declaration.
const testVersion = 4 // find the value in gigasecond_test.go

//AddGigasecond add a gigasecond to the time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(1e9) * time.Second)
}
