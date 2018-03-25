package clock

import (
	"fmt"
	)

// A Clock represents a minute span within a day
type Clock struct {
	hours, minutes int
}

// New creates a Clock taking two ints for hour and minutes. Any int input is allowed, including negative values
// The values will be normalized to Military Time Chart ranges:  0 <= hrs < 24, 0 <= mins < 60
func New(hours, minutes int) Clock {
	hours = hours + (minutes / 60)
	minutes = minutes % 60
	if minutes < 0 {
		minutes += 60
		hours--
	}
	hours = hours % 24
	if hours < 0 {
		hours += 24
	}
	return Clock{hours, minutes}
}

// String allows Clock to be a stringer
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

// Add returns a new valid Clock instance adding `a` minutes (allows negatives) to this Clock
func (c Clock) Add(a int) Clock {
	return New(c.hours, c.minutes+a)
}

// Subtract returns a new valid Clock instance subtracting `a` minutes (if negative behaves like Add with the absolute value of `a`) to this Clock
func (c Clock) Subtract(a int) Clock {
	return New(c.hours, c.minutes-a)
}
