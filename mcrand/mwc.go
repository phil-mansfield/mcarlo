package mcrand

import (
	"math"
)

const (
	phi uint32 = 0x9e3779b9
	a uint64 = 18782
	r uint32 = 0xfffffffe
)

type multiplyWithCarryGenerator struct {
	Q [4096]uint32
	c uint32
	i int32
}

func (gen *multiplyWithCarryGenerator) Init(seed int64) {
	x := uint32(seed)
	gen.Q[0] = x
	gen.Q[1] = x + phi
	gen.Q[2] = x + phi + phi

	for gen.i = 3; gen.i < 4096; gen.i++ {
		gen.Q[gen.i] = gen.Q[gen.i - 3] ^ gen.Q[gen.i - 2] ^ phi ^ uint32(gen.i)
	}

	gen.c = 362463
}

func (gen *multiplyWithCarryGenerator) Next() float64 {
	gen.i = (gen.i + 1) & 4095

	t := a * uint64(gen.Q[gen.i]) + uint64(gen.c)
	gen.c = uint32(t >> 32)
	x := uint32(t + uint64(gen.c))
	if x < gen.c {
		x++
		gen.c++
	}
	gen.Q[gen.i] = r - x
	return float64(gen.Q[gen.i]) / float64(math.MaxUint32)
}

func (gen *multiplyWithCarryGenerator) NextN(n int) []float64 {
	xs := make([]float64, n)
	for i := 0; i < n; i++ {
		xs[i] = gen.Next()
	}
	return xs
}
