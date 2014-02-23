package mcrand

import (
	"testing"
)

func BenchmarkTauswortheNext(b *testing.B) {
	rand := NewTimeSeed(Tausworthe)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rand.Uniform(0, 2)
	}
}

func BenchmarkGoRandNext(b *testing.B) {
	rand := NewTimeSeed(GoRand)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rand.Uniform(0, 2)
	}
}

func BenchmarkXorshiftNext(b *testing.B) {
	rand := NewTimeSeed(Xorshift)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rand.Uniform(0, 2)
	}
}

func BenchmarkMultiplyWithCarryNext(b *testing.B) {
	rand := NewTimeSeed(MultiplyWithCarry)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rand.Uniform(0, 2)
	}
}
