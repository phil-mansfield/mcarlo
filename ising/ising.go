package ising

type IsingModel interface {
	Init(width uint32)

	SetMag(mag float64)

	Sites() uint32

	Flip(site uint32)

	Energy(site uint32) float64
	Magnetization(site uint32) float64
}

type IsingModelType uint8
const (
	Grid2DType IsingModelType = iota
)

func New(imt IsingModelType, width uint32) IsingModel {
	var model IsingModel
	switch(imt) {
	case Grid2DType:
		model = new(Grid2D)
	default:
		panic("Unrecognized IsingModelType")
	}

	model.Init(width)
	return model
}
