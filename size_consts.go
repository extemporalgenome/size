package size

// Binary byte multiples
const (
	_ Size = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
)

// Decimal byte multiples
const (
	KB Size = 1000
	MB      = KB * 1000
	GB      = MB * 1000
	TB      = GB * 1000
	PB      = TB * 1000
	EB      = PB * 1000
)

// Binary bit multiples, converted to bytes
const (
	_ Size = 1 << (10 * iota) / 8
	Kibit
	Mibit
	Gibit
	Tibit
	Pibit
	Eibit
)

// Decimal bit multiples, converted to bytes
const (
	Kbit Size = 1000 / 8
	Mbit      = Kbit * 1000
	Gbit      = Mbit * 1000
	Tbit      = Gbit * 1000
	Pbit      = Tbit * 1000
	Ebit      = Pbit * 1000
)
