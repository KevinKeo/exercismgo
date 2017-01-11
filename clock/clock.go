/*
Package clock which provides methods to operate over a simple clock of hours and minutes
*/
package clock

import (
	"strconv"
)

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

/*
Struct Clock which contain hour and minute
*/
type Clock struct {
	hour, minute int
}

/*
New create a new Clock and take care of negative hours and mintues
*/
func New(hour, minute int) Clock {
	hour += minute / 60
	minute = minute % 60
	if minute < 0 {
		minute += 60
		hour -= 1
	}
	hour = hour % 24
	if hour < 0 {
		hour += 24
	}
	return Clock{hour, minute}

}

func (c Clock) String() string {
	time := ""
	if c.hour < 10 {
		time += "0"
	}
	time += strconv.Itoa(c.hour) + ":"
	if c.minute < 10 {
		time += "0"
	}
	time += strconv.Itoa(c.minute)
	return time
}

func (c Clock) Add(minutes int) Clock {
	newHours := c.hour + (minutes / 60)
	newMinutes := minutes%60 + c.minute
	return New(newHours, newMinutes)
}
