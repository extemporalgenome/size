// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package size provides byte/bit size constants and utilities.
//
// All values are stored as byte quantities (octets). Bit quantities which
// are not an exact multiple of eight bits are therefore not precisely
// representable, though such cases are exceptionally rare.
package size
