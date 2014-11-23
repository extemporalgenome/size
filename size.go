package size

import "time"

// Size represents a quantity of bytes
type Size int64

// Rate calculates the byte-rate needed to reach the input
// size over the given duration.
func (s Size) Rate(dur time.Duration) Rate {
	return Rate(s / Size(dur/time.Second))
}
