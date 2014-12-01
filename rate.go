package size

import (
	"math"
	"time"
)

// Rate represents a quantity of bytes per second.
type Rate int64

// Size returns the number of bytes transferrable given
// the input rate and duration.
func (r Rate) Size(dur time.Duration) Size {
	return Size(r * Rate(dur/time.Second))
}

// NBytes returns the quantity of bytes/second that the Rate represents.
func (r Rate) NBytes() int64 {
	return int64(r)
}

// NBits returns the quantity of bits/second that the Rate represents.
// If the input Rate is not in the non-inclusive range (-1 EiB/s, 1 EiB/s),
// rather than wrap around on overflow, the result will be truncated
// to the minimum or maximum int64 value; this condition can be checked
// by passing the result to IsBitOverflow.
func (r Rate) NBits() int64 {
	switch {
	case r >= EiBs:
		return math.MaxInt64
	case r <= -EiBs:
		return math.MinInt64
	}
	return int64(r * 8)
}
