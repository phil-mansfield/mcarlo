package mcrand

type Generator interface {
	Init(seed int64) // Certain uber-secure algorithms ignore the seed.
	Next() float64
	NextN(n int) []float64
}

type GeneratorType uint8
const (
	Xorshift GeneratorType = iota
	GoRand
	GoCrypto
	MultiplyWithCarry
	BlumBlum
)

func New(gt GeneratorType, seed int64) Generator {
	var gen Generator

	switch(gt) {
	case Xorshift:
		gen = new(xorshiftGenerator)
	case GoRand:
		gen = new(goRandGenerator)
	case GoCrypto:
		gen = new(goCryptoGenerator)
	case BlumBlum:
		gen = new(blumBlumGenerator)
	case MultiplyWithCarry:
		gen = new(multiplyWithCarryGenerator)
	default:
		panic("Unrecognized GeneratorType")
	}

	gen.Init(seed)
	return gen
}
