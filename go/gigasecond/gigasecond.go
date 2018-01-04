package gigasecond

import "time"

// AddGigasecond adds a gigasecond (10^9 seconds) to the time value passed
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
