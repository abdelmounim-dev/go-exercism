// Package gigasecond exports a function AddGigasecond that adds one gigasecond (1e9 seconds) to a given date
package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

// AddGigasecond: adds one gigasecond (1e9 seconds) to a given date
func AddGigasecond(t time.Time) time.Time {
	gigasecond := time.Duration(1e9 * float64(time.Second))

	return t.Add(gigasecond)
}
