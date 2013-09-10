package mcrand

import (
	"math"
)

type xorshiftGenerator struct {
	w, x, y, z uint32
}

func (gen *xorshiftGenerator) Init(seed int64) {
	gen.x = 123456789
	gen.y = 362436069
	gen.z = 521288629
	gen.w = uint32(seed)
}

func (gen *xorshiftGenerator) Next() float64 {
	t := gen.x ^ (gen.x << 11)
	gen.x, gen.y, gen.z = gen.y, gen.z, gen.w
	gen.w = gen.w ^ (gen.w >> 19) ^ (t ^ (t >> 8))
	return float64(gen.w) / float64(math.MaxUint32)
}

func (gen *xorshiftGenerator) NextN(n int) []float64 {
	xs := make([]float64, n)
	w, x, y, z := gen.w, gen.x, gen.y, gen.z
	for i := 0; i < n; i++ {
		t := x ^ (x << 11)
		x, y, z = y, z, w
		w = w ^ (w >> 19) ^ (t ^ (t >> 8))
		xs[i] = float64(w) / float64(math.MaxUint32)
	}
	gen.w, gen.x, gen.y, gen.z = w, x, y, z
	return xs
}
