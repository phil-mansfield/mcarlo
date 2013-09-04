package ising

type Grid2D struct {
	cells []bool
	mag float64
	width uint32
}

func (g *Grid2D) Init(width uint32, mag float64) {
	g.cells = make([]uint32, width * width)
	g.mag = mag
	g.width = width
}

func (g *Grid2D) Sites() uint32 { return len(g.cells) }

func (g *Grid2D) Flip(uint32 site) { cells[site] = !cells[site] }

func (g *Gird2D) FlipEnergy(uint32 site) float64 { 
	return g.Energy(site) - g.Magnetization(site) * g.mag
}

func 

func (g *Grid2D) Energy(uint32 site) float64 {
	x := site % g.width
	y := site / g.width

	sqr := g.width * g.width

	mc := g.Magnetization(site)
	mr := g.Magnetization()
	ml := g.Magnetization()
	md := g.Magnetization()
	mu := g.Magnetization((sqr + site) % sqr)

	retunr + g.Magnetization(site) * g.mag
}


func (g *Grid2D) Magnetization(uint32 site) float64 {
	if g.cells[site] {
		return 1.0
	}
	return -1.0
}
