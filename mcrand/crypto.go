package mcrand

import (
	"crypto/rand"
	"math"
	"math/big"
)

type goCryptoGenerator struct { }

func (gen *goCryptoGenerator) Init(seed int64) { }

func (gen *goCryptoGenerator) Next() float64 {
	num, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic(err)
	}
	return float64(num.Int64()) / float64(math.MaxInt64)
}

func (gen *goCryptoGenerator) NextN(n int) []float64 {
	xs := make([]float64, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
		if err != nil {
			panic(err)
		}
		xs[i] = float64(num.Int64()) / float64(math.MaxInt64)
	}
	return xs
}
