package mcrand

import (
	"crypto/rand"
	"math"
	"math/big"
)

const (
	nBits = 128
)

type blumBlumGenerator struct {
	N, state *big.Int
}

func newPrime(bits int) *big.Int {
	res := big.NewInt(0)

	for {
		p, err := rand.Prime(rand.Reader, bits)
		if err != nil {
			panic(err)
		}

		res.Mod(p, big.NewInt(4))
		if res.Int64() == 3 {
			return p
		}
	}
}

func bigEq(x, y *big.Int) bool {
	res := big.NewInt(0)
	res.Sub(x, y)
	return res.Int64() == 0
}

func (gen *blumBlumGenerator) Init(seed int64) {
	p, q := newPrime(nBits / 2), newPrime(nBits / 2)
	for bigEq(p, q) {
		q = newPrime(nBits / 2)
	}

	gen.N = big.NewInt(0)
	gen.N.Mul(p, q)

	gen.state = newPrime(nBits)
}

func (gen *blumBlumGenerator) Next() float64 {
	var result int64 = 0
	res := big.NewInt(0)
	for i := 0; i < 64; i++ {
		res.Mul(gen.state, gen.state)
		gen.state.Mod(res, gen.N)
		result = (result << 1) | (res.And(gen.state, big.NewInt(1)).Int64())
	}
	return float64(result) / float64(math.MaxInt64)
}

func (gen *blumBlumGenerator) NextN(n int) []float64 {
	xs := make([]float64, n)
	for i := 0; i < n; i++ {
		xs[i] = gen.Next()
	}
	return xs
}
