package mcrand

import (
	"testing"
)

func BenchmarkBlumBlumArray(b *testing.B) {
	rand := New(BlumBlum, 0)

	b.ReportAllocs()
	b.ResetTimer()

	xs := rand.NextN(b.N)

	sum := 0.0
	for i := 0; i < b.N; i++ {
		sum += xs[i]
	}
}

func BenchmarkBlumBlumNext(b *testing.B) {
	rand := New(BlumBlum, 0)

	b.ReportAllocs()
	b.ResetTimer()

	sum := 0.0
	for i := 0; i < b.N; i++ {
		sum += rand.Next()
	}
}

func BenchmarkGoRandArray(b *testing.B) {
	rand := New(GoRand, 0)

	b.ReportAllocs()
	b.ResetTimer()

	xs := rand.NextN(b.N)

	sum := 0.0
	for i := 0; i < b.N; i++ {
		sum += xs[i]
	}
}

func BenchmarkGoRandNext(b *testing.B) {
	rand := New(GoRand, 0)

	b.ReportAllocs()
	b.ResetTimer()

	sum := 0.0
	for i := 0; i < b.N; i++ {
		sum += rand.Next()
	}
}

func BenchmarkGoCryptoArray(b *testing.B) {
	rand := New(GoCrypto, 0)

	b.ReportAllocs()
	b.ResetTimer()

	xs := rand.NextN(b.N)

	sum := 0.0
	for i := 0; i < b.N; i++ {
		sum += xs[i]
	}
}

func BenchmarkGoCryptoNext(b *testing.B) {
	rand := New(GoCrypto, 0)

	b.ReportAllocs()
	b.ResetTimer()

	sum := 0.0
	for i := 0; i < b.N; i++ {
		sum += rand.Next()
	}
}

func BenchmarkXorshiftArray(b *testing.B) {
	rand := New(Xorshift, 0)

	b.ReportAllocs()
	b.ResetTimer()

	xs := rand.NextN(b.N)

	sum := 0.0
	for i := 0; i < b.N; i++ {
		sum += xs[i]
	}
}

func BenchmarkXorshiftNext(b *testing.B) {
	rand := New(Xorshift, 0)

	b.ReportAllocs()
	b.ResetTimer()

	sum := 0.0
	for i := 0; i < b.N; i++ {
		sum += rand.Next()
	}
}

func BenchmarkMultiplyWithCarryArray(b *testing.B) {
	rand := New(MultiplyWithCarry, 0)

	b.ReportAllocs()
	b.ResetTimer()

	xs := rand.NextN(b.N)

	sum := 0.0
	for i := 0; i < b.N; i++ {
		sum += xs[i]
	}
}

func BenchmarkMultiplyWithCarryNext(b *testing.B) {
	rand := New(MultiplyWithCarry, 0)

	b.ReportAllocs()
	b.ResetTimer()

	sum := 0.0
	for i := 0; i < b.N; i++ {
		sum += rand.Next()
	}
}
