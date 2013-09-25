package ising

import (
	"math"
	
	"github.com/phil-mansfield/mcarlo/ising/hist"
	"github.com/phil-mansfield/mcarlo/mcrand"
)

type Grid2D struct {
	cells []bool
	hist *hist.Histogram
	mag, totalEnergy, totalMag, minEnergy, histStep, temp float64
	width, modMask, sqrModMask uint32

	precalcTable []float64
}

func (g *Grid2D) Init(widthPow uint32) {
	g.width = 1 << widthPow
	g.modMask = g.width - 1
	g.sqrModMask = g.width * g.width - 1

	g.cells = make([]bool, g.width * g.width)
	for i := 0; i < len(g.cells); i++ {
		g.cells[i] = true
	}

	g.totalMag = float64(g.width)
	g.totalEnergy = - float64(2 * g.width * g.width)

	g.hist = hist.New(1 + int(g.width * g.width), g.totalEnergy, 4.0)
}

func (g *Grid2D) SetMag(mag float64) {
	panic("For optimization reasons, this is no longer supported.")

	g.totalEnergy -= g.TotalMagnetization() * g.mag
	g.mag = mag
	g.totalMag += g.TotalMagnetization() * g.mag
}

func (g *Grid2D) Sites() uint32 { 
	return uint32(len(g.cells))
}

func (g *Grid2D) SetTemp(temp float64) {
	if temp == g.temp && len(g.precalcTable) != 0 {
		return
	}

	g.temp = temp

	g.precalcTable = make([]float64, 4 + 1)
	for i := 0; i < len(g.precalcTable); i++ {
		g.precalcTable[i] = math.Exp(-2 * float64(i) / temp)
	}
}

func (g *Grid2D) TryToFlip(site uint32, gen mcrand.Generator) bool { 
	E := g.intEnergy(site)
	dE := -2.0 * float64(E)

	if dE < 0 || gen.Next() < g.precalcTable[-E] {
		g.cells[site] = !g.cells[site]
		g.totalEnergy += dE
		if g.cells[site] {
			g.totalMag += 2.0
		} else {
			g.totalMag -= 2.0
		}
		return true
	}
	return false
}

func (g *Grid2D) intMag(site uint32) int {
	if g.cells[site] {
		return 1
	}
	return -1
}

func (g *Grid2D) intEnergy(site uint32) int {
	var bitCrap uint32 = 0
	upperBits := site & ((bitCrap - 1) ^ g.modMask)
	mr := g.intMag(upperBits | ((site + 1) & g.modMask))
	ml := g.intMag(upperBits | ((site - 1) & g.modMask))
	md := g.intMag((site + g.width) & g.sqrModMask)
	mu := g.intMag((site - g.width) & g.sqrModMask)
	m := g.intMag(site)

	return -m * (mr + ml + md + mu)
}

func (g *Grid2D) Energy(site uint32) float64 {
	return float64(g.intMag(site))
}

func (g *Grid2D) Magnetization(site uint32) float64 {
	if g.cells[site] {
		return 1.0
	}
	return -1.0
}

func (g *Grid2D) TotalEnergy() float64 {
	return g.totalEnergy
}

func (g *Grid2D) TotalMagnetization() float64 {
	return g.totalMag
}

func (g *Grid2D) ResetHistogram() {
	for i := 0; i < len(g.hist.Hist); i++ {
		g.hist.Hist[i] = 0
	}
}

func (g *Grid2D) EnergyHistogram() *hist.Histogram {
	return g.hist
}

func (g *Grid2D) Print() {
	for y := 0; y < int(g.width); y++ {
		for x := 0; x < int(g.width); x++ {
			if g.cells[x + y * int(g.width)] {
				print("+")
			} else {
				print("-")
			}
		}
		println()
	}
}
