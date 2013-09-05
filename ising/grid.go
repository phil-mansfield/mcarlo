package ising

import (
	"math/rand"
)

type Grid2D struct {
	cells []bool
	mag float64
	width uint32
}

func (g *Grid2D) Init(width uint32) {
	g.cells = make([]bool, width * width)
	for i := 0; i < len(g.cells); i++ {
		g.cells[i] = rand.Intn(2) == 1
	}

	g.width = width
}

func (g *Grid2D) SetMag(mag float64) {
	g.mag = mag
}

func (g *Grid2D) Sites() uint32 { 
	return uint32(len(g.cells))
}

func (g *Grid2D) Flip(site uint32) { 
	g.cells[site] = !g.cells[site]
}

func (g *Grid2D) up(site uint32) uint32 {
	return (uint32(len(g.cells)) + site - g.width) % uint32(len(g.cells))
}

func (g *Grid2D) down(site uint32) uint32 {
	return (site + g.width) % uint32(len(g.cells))
}

func (g *Grid2D) right(site uint32) uint32 {
	if site % g.width == g.width - 1 {
		return site - (g.width - 1)
	}
	return site + 1
}

func (g *Grid2D) left(site uint32) uint32 {
	if site % g.width == 0 {
		return site + (g.width - 1)
	}
	return site - 1
}

func (g *Grid2D) Energy(site uint32) float64 {
	mc := g.Magnetization(site)
	mr := g.Magnetization(g.right(site))
	ml := g.Magnetization(g.left(site))
	md := g.Magnetization(g.down(site))
	mu := g.Magnetization(g.up(site))

	return -mc * (mr + ml + md + mu + g.mag)
}


func (g *Grid2D) Magnetization(site uint32) float64 {
	if g.cells[site] {
		return 1.0
	}
	return -1.0
}

func (g *Grid2D) Print() {
	panic("Meow")
}
