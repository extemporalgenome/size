package size

// Binary byterate multiples.
const (
	_ Rate = 1 << (10 * iota)
	KiBs
	MiBs
	GiBs
	TiBs
	PiBs
	EiBs
)

// Decimal byterate multiples.
const (
	KBs Rate = 1000
	MBs      = KBs * 1000
	GBs      = MBs * 1000
	TBs      = GBs * 1000
	PBs      = TBs * 1000
	EBs      = PBs * 1000
)

// Binary bitrate multiples, expressed in bytes/sec.
const (
	_ Rate = 1 << (10 * iota) / 8
	Kibits
	Mibits
	Gibits
	Tibits
	Pibits
	Eibits
)

// Decimal bitrate multiples, expressed in bytes/sec.
const (
	Kbits Rate = 1000 / 8
	Mbits      = Kbits * 1000
	Gbits      = Mbits * 1000
	Tbits      = Gbits * 1000
	Pbits      = Tbits * 1000
	Ebits      = Pbits * 1000
)
