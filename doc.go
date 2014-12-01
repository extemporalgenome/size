// Package size provides byte/bit size and rate constants.
//
// All Size and Rate values are internally stored as quanties
// of bytes. This package considers a byte to be precisely
// one octet (8 bits) in size. Bit quantities that are not an
// exact multiple of eight are therefore not accurately
// representable, though such cases are exeptionally rare.
//
// Any value in the range (-1 EiB, 1 EiB) [non-inclusive] is
// supported when using the NBits methods. When conversion to
// bit quantities is not needed,
package size
