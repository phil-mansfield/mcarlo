package ising

import (
	"fmt"

	"github.com/phil-mansfield/mcarlo/ising/hist"
	"github.com/phil-mansfield/mcarlo/mcrand"
)

type IsingModel interface {
	Init(widthPow uint32) // All cells are initially set to true.

	Sites() uint32

	SetTemp(temp float64)
	TryToFlip(site uint32, gen mcrand.Generator) bool

	ResetHistogram()
	EnergyHistogram() *hist.Histogram
	
	Energy(site uint32) float64
	Magnetization(site uint32) float64
	TotalEnergy() float64
	TotalMagnetization() float64
	
	Print()
}

type SweepInfo struct {
	Im IsingModel
	St SweepType
	Gen mcrand.Generator	
	Temp float64
}

type SweepStats struct {
	EHistory, MagHistory []float64
	E, Mag, C, X float64
}

type IsingModelType uint8
const (
	Grid2DType IsingModelType = iota
)

type SweepType uint8
const (
	RandomSweep SweepType = iota
	IterSweep
	CheckerSweep
)

func New(imt IsingModelType, widthPow uint32) IsingModel {
	var model IsingModel
	switch(imt) {
	case Grid2DType:
		model = new(Grid2D)
	default:
		panic("Unrecognized IsingModelType")
	}

	model.Init(widthPow)
	return model
}

func randomSweep(im IsingModel, g mcrand.Generator, sweeps int) {
	for j := 0; j < sweeps; j++ {
		for i := 0; i < int(im.Sites()); i++ {
			site := uint32(float64(im.Sites()) * g.Next())
			im.TryToFlip(site, g)
		}
		im.EnergyHistogram().AddFloat(im.TotalEnergy())
	}
}

func iterSweep(im IsingModel, g mcrand.Generator, sweeps int) {
	for s := 0; s < sweeps; s++ {
		var site uint32 = 0
		for ; site < im.Sites(); site++ {
			im.TryToFlip(site, g)
		}
		im.EnergyHistogram().AddFloat(im.TotalEnergy())
	}
}

// This is suboptimal.
func sqrtBase2(n uint32) uint32 {
	var i uint32 = 1
	for ; i <= 16; i++ {
		n = n >> 2
		if n == 0 {
			return 1 << (i * 2)
		}
	}
	panic(fmt.Sprintf("%d is not a square of a power of two", n))
}

func checkerSweep(im IsingModel, g mcrand.Generator, sweeps int) {
	for s := 0; s < sweeps; s++ {
		widthMask := sqrtBase2(im.Sites())

		offset := -1
		var site uint32 = 0
		for ; site < uint32(im.Sites()); site += 2 {
			if site & widthMask <= 1 && site > 1 {
				offset *= -1
				site = uint32(int(site) + offset)
			}
			im.TryToFlip(site, g)
		}

		site = 1
		offset = 1
		for ; site < uint32(im.Sites()); site += 2 {
			if site & widthMask <= 1 && site > 1 {
				offset *= -1
				site = uint32(int(site) + offset)
			}
			im.TryToFlip(site, g)
		}

		im.EnergyHistogram().AddFloat(im.TotalEnergy())
	}
}

func Sweep(info *SweepInfo, sweeps int) {
	info.Im.SetTemp(info.Temp)
	switch(info.St) {
	case RandomSweep:
		randomSweep(info.Im, info.Gen, sweeps)
		return
	case IterSweep:
		iterSweep(info.Im, info.Gen, sweeps)
		return
	case CheckerSweep:
		checkerSweep(info.Im, info.Gen, sweeps)
		return
	}

	panic("What are you doing?")
}

func variance(xs []float64) float64 {
	sqrSum := 0.0
	sum := 0.0
	for _, x := range xs {
		sum += x
		sqrSum += x * x
	}
	return sqrSum / float64(len(xs)) - sum * sum / float64(len(xs) * len(xs))
}

func sum(xs []float64) float64 {
	total := 0.0
	for _, x := range xs {
		total += x
	}
	return total
}

func avg(xs []float64) float64 {
	return sum(xs) / float64(len(xs))
}

func SweepStatistics(info *SweepInfo, initialSweeps, trials int) *SweepStats {
	// Reach equilibrium
	Sweep(info, initialSweeps)
	info.Im.ResetHistogram()

	stats := &SweepStats{
		make([]float64, 0, trials),
		make([]float64, 0, trials),
		0.0, 0.0, 0.0, 0.0,
	}

	N := info.Im.Sites()
	for i := 0; i < trials; i++ {
		Sweep(info, 1)

		E, mag := 0.0, 0.0
		var site uint32 = 0
		for ; site < N; site++ {
			E += info.Im.Energy(site)
			mag += info.Im.Magnetization(site)
		}
		
		E, mag = E / float64(N), mag / float64(N)
		stats.EHistory = append(stats.EHistory, E)
		stats.MagHistory = append(stats.MagHistory, mag)
	}

	stats.E, stats.Mag = avg(stats.EHistory), avg(stats.MagHistory)
	stats.C = float64(N) * variance(stats.EHistory) / (info.Temp * info.Temp)
	stats.X = float64(N) * variance(stats.MagHistory) / info.Temp

	return stats
}

// This is not highly optimized.
func Tau(xs []float64) (fs []float64, tau float64) {
	fs = make([]float64, len(xs))

	averageX := avg(xs)
	varianceX := variance(xs)
	sqrAvg := averageX * averageX

	for t := 0; t < len(xs); t++ {
		corr := 0.0
		for tp := 0; tp < len(xs) - t; tp++ {
			corr += xs[tp] * xs[tp + t]
		}
		corr /= float64(len(xs) - t)

		fs[t] = (corr - sqrAvg) / varianceX
		if fs[t] < 0 {
			fs = fs[0: t]
			break
		}
	}

	return fs, sum(fs)
}

func AvgSweepStats(stats []*SweepStats) *SweepStats {
	avgE, avgMag, avgC, avgX := 0.0, 0.0, 0.0, 0.0
	for _, stat := range stats {
		avgE += stat.E
		avgMag += stat.Mag
		avgC += stat.C
		avgX += stat.X
	}
	n := float64(len(stats))
	avgE, avgMag, avgC, avgX = avgE / n, avgMag / n, avgC / n, avgX / n
	return &SweepStats{ nil, nil, avgE, avgMag, avgC, avgX, }
}

func GaussianMax(xs, ys []float64) (maxX, maxY float64) {
	panic("I HATE YOU SO MUCH, GO")
}
