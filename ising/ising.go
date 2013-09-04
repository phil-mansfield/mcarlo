package main

type IsingModel interface {
	Init(width uint32)

	Sites() uint32

	Flip(uint32 site)
	FlipEnergy(uint32 site) float64

	Energy(uint32 site) float64
	Magnetization(uint32 site) float64
}

type IsingModelType uint8
const (
	Grid2D IsingModelType = iota
)

func New(imt IsingModelType, width uint32) *IsingModel {
	var model *IsingModel
	switch(imt) {
	case Grid2D:
		model = new(Grid2D)
	default:
		panic("Unrecognized IsingModelType")
	}

	model.Init(width)
	return model
}
