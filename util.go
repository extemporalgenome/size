// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package size

import "math"

// IsBitOverflow returns true if nbits represents a special overflow value.
// nbits is the result of calling the NBits method on a Size or Rate value.
func IsBitOverflow(nbits int64) bool {
	return nbits == math.MinInt64 || nbits == math.MaxInt64
}
