package mcrand

import (
	"math/rand"
)

type goRandGenerator struct {
	r *rand.Rand
}

func (gen *goRandGenerator) Init(seed int64) {
	src := rand.NewSource(seed)
	gen.r = rand.New(src)
}

func (gen *goRandGenerator) Next() float64 {
	return gen.r.Float64()
}

func (gen *goRandGenerator) NextN(n int) []float64 {
	xs := make([]float64, n)
	for i := 0; i < n; i++ {
		xs[i] = gen.r.Float64()
	}
	return xs
}
