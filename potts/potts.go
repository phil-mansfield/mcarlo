package potts

import (
	"math/rand"
)

type GridShape uint8
const (
	Grid2D GridShape = iota
	Grid3D
	Hexagon
)

func generateGrid2D(prob float64, w int) (us, vs []uint32) {
	width := uint32(w)
	height := width

	us = make([]uint32, 0, height * width * 2)
	vs = make([]uint32, 0, height * width * 2)

	var x, y uint32
	for y = 0; y < height; y++ {
		for x = 0; x < width; x++ {
			i := x + y * width

			right := ((x + 1) % width) + y * width
			down := x + ((y + 1) % height) * width
			
			if rand.Float64() < prob {
				us = append(us, i)
				vs = append(vs, right)
			}

			if rand.Float64() < prob {
				us = append(us, i)
				vs = append(vs, down)
			}
		}
	}
}

func GenerateEdges(shape GridShape, prob float64, width int) (us, vs []uint32) {
	switch(shape) {
	case Grid2D:
		return generateGrid2D(prob, width)
	case Grid3D:
		panic("potts.Grid3D is not yet implemented.")
	case Hexagon:
		panic("potts.Hexagon is not yet implemented.")
	}
	panic("Unrecognized potts.GridShape type.")
}
