package size

import (
	"math"
	"time"
)

// Size represents a quantity of bytes.
type Size int64

// Rate returns the byte-rate required to transfer the input size over the
// given duration.
func (s Size) Rate(dur time.Duration) Rate {
	return Rate(s / Size(dur/time.Second))
}

// Duration returns the length of time needed to complete a transfer of the
// given Size at the given Rate.
func (s Size) Duration(r Rate) time.Duration {
	sPrime := time.Duration(s)
	rPrime := time.Duration(r)
	if sPrime > math.MaxInt64/time.Second {
		// avoid overflow extreme case by using less-accurate expression
		return sPrime / rPrime * time.Second
	}
	return sPrime * time.Second / rPrime
}

// NBytes returns the quantity of bytes that the Size represents.
func (s Size) NBytes() int64 {
	return int64(s)
}

// NBits returns the quantity of bits that the Size represents. If the
// input Size is not in the non-inclusive range (-1 EiB, 1 EiB), rather
// than wrap around on overflow, the result will be truncated to the
// minimum or maximum int64 value; this condition can be checked by passing
// the result to IsBitOverflow.
func (s Size) NBits() int64 {
	switch {
	case s >= EiB:
		return math.MaxInt64
	case s <= -EiB:
		return math.MinInt64
	}
	return int64(s * 8)
}
