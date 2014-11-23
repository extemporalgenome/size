package size

import "time"

// Rate represents a quantity of bytes per second
type Rate int64

// Size calculates the number of bytes transferrable given
// the input rate and duration.
func (r Rate) Size(dur time.Duration) Size {
	return Size(r * Rate(dur/time.Second))
}
