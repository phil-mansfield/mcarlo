package ising

import (
	"testing"

	"github.com/phil-mansfield/mcarlo/mcrand"
)

const (
	testTemp = 2.5
)

// RandomSweep

func BenchmarkRandomSweep4(b *testing.B) {
	im := New(Grid2DType, 2)
	b.ResetTimer()
	info := &SweepInfo{ 
		im,
		RandomSweep,
		mcrand.New(mcrand.GoRand, 0),
		testTemp,
	}
	Sweep(info, b.N)
}

func BenchmarkRandomSweep8(b *testing.B) {
	im := New(Grid2DType, 3)
	b.ResetTimer()
	info := &SweepInfo{ 
		im,
		RandomSweep,
		mcrand.New(mcrand.GoRand, 0),
		testTemp,
	}
	Sweep(info, b.N)
}

func BenchmarkRandomSweep16(b *testing.B) {
	im := New(Grid2DType, 4)
	b.ResetTimer()
	info := &SweepInfo{ 
		im,
		RandomSweep,
		mcrand.New(mcrand.GoRand, 0),
		testTemp,
	}
	Sweep(info, b.N)
}

func BenchmarkRandomSweep32(b *testing.B) {
	im := New(Grid2DType, 5)
	b.ResetTimer()
	info := &SweepInfo{ 
		im,
		RandomSweep,
		mcrand.New(mcrand.GoRand, 0),
		testTemp,
	}
	Sweep(info, b.N)
}

func BenchmarkRandomSweep64(b *testing.B) {
	im := New(Grid2DType, 6)
	b.ResetTimer()
	info := &SweepInfo{ 
		im,
		RandomSweep,
		mcrand.New(mcrand.GoRand, 0),
		testTemp,
	}
	Sweep(info, b.N)
}

// IterSweep

func BenchmarkIterateSweep4(b *testing.B) {
	im := New(Grid2DType, 2)
	b.ResetTimer()
	info := &SweepInfo{ 
		im,
		IterSweep,
		mcrand.New(mcrand.GoRand, 0),
		testTemp,
	}
	Sweep(info, b.N)
}

func BenchmarkIterateSweep8(b *testing.B) {
	im := New(Grid2DType, 3)
	b.ResetTimer()
	info := &SweepInfo{ 
		im,
		IterSweep,
		mcrand.New(mcrand.GoRand, 0),
		testTemp,
	}
	Sweep(info, b.N)
}

func BenchmarkIterateSweep16(b *testing.B) {
	im := New(Grid2DType, 4)
	b.ResetTimer()
	info := &SweepInfo{ 
		im,
		IterSweep,
		mcrand.New(mcrand.GoRand, 0),
		testTemp,
	}
	Sweep(info, b.N)
}

func BenchmarkIterateSweep32(b *testing.B) {
	im := New(Grid2DType, 5)
	b.ResetTimer()
	info := &SweepInfo{ 
		im,
		IterSweep,
		mcrand.New(mcrand.GoRand, 0),
		testTemp,
	}
	Sweep(info, b.N)
}

func BenchmarkIterateSweep64(b *testing.B) {
	im := New(Grid2DType, 6)
	b.ResetTimer()
	info := &SweepInfo{ 
		im,
		IterSweep,
		mcrand.New(mcrand.GoRand, 0),
		testTemp,
	}
	Sweep(info, b.N)
}

// CheckerSweep

func BenchmarkCheckerSweep4(b *testing.B) {
	im := New(Grid2DType, 2)
	b.ResetTimer()
	info := &SweepInfo{ 
		im,
		CheckerSweep,
		mcrand.New(mcrand.GoRand, 0),
		testTemp,
	}
	Sweep(info, b.N)
}

func BenchmarkCheckerSweep8(b *testing.B) {
	im := New(Grid2DType, 3)
	b.ResetTimer()
	info := &SweepInfo{ 
		im,
		CheckerSweep,
		mcrand.New(mcrand.GoRand, 0),
		testTemp,
	}
	Sweep(info, b.N)
}

func BenchmarkCheckerSweep16(b *testing.B) {
	im := New(Grid2DType, 4)
	b.ResetTimer()
	info := &SweepInfo{ 
		im,
		CheckerSweep,
		mcrand.New(mcrand.GoRand, 0),
		testTemp,
	}
	Sweep(info, b.N)
}

func BenchmarkCheckerSweep32(b *testing.B) {
	im := New(Grid2DType, 5)
	b.ResetTimer()
	info := &SweepInfo{ 
		im,
		CheckerSweep,
		mcrand.New(mcrand.GoRand, 0),
		testTemp,
	}
	Sweep(info, b.N)
}

func BenchmarkCheckerSweep64(b *testing.B) {
	im := New(Grid2DType, 6)
	b.ResetTimer()
	info := &SweepInfo{ 
		im,
		CheckerSweep,
		mcrand.New(mcrand.GoRand, 0),
		testTemp,
	}
	Sweep(info, b.N)
}

// RNG conparisons

func BenchmarkGoRandGenerator(b *testing.B) {
	im := New(Grid2DType, 5)
	b.ResetTimer()
	info := &SweepInfo{ 
		im,
		CheckerSweep,
		mcrand.New(mcrand.GoRand, 0),
		testTemp,
	}
	Sweep(info, b.N)
}

func BenchmarkGoCryptoGenerator(b *testing.B) {
	im := New(Grid2DType, 5)
	b.ResetTimer()
	info := &SweepInfo{ 
		im,
		CheckerSweep,
		mcrand.New(mcrand.GoCrypto, 0),
		testTemp,
	}
	Sweep(info, b.N)
}

func BenchmarkXorshiftGenerator(b *testing.B) {
	im := New(Grid2DType, 5)
	b.ResetTimer()
	info := &SweepInfo{ 
		im,
		CheckerSweep,
		mcrand.New(mcrand.Xorshift, 0),
		testTemp,
	}
	Sweep(info, b.N)
}

func BenchmarkMultiplyWithCarryGenerator(b *testing.B) {
	im := New(Grid2DType, 5)
	b.ResetTimer()
	info := &SweepInfo{ 
		im,
		CheckerSweep,
		mcrand.New(mcrand.MultiplyWithCarry, 0),
		testTemp,
	}
	Sweep(info, b.N)
}