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

func Equilibriate(im IsingModel, steps int) {
	return
}

func AverageEnergy(im IsingModel) float64 {
	energySum := 0.0

	var i uint32
	for i = 0; i < im.Sites(); i++ {
		energySum += im.Energy(i)
	}
	return energySum / float64(2 * im.Sites())
}

func AverageMag(im IsingModel) float64 {
	magSum := 0.0

	var i uint32
	for i = 0; i < im.Sites(); i++ {
		magSum += im.Magnetization(i)
	}
	return magSum / float64(im.Sites())
}

func SpecificHeat(im IsingModel) float64 {
	sqrSum := 0.0

	var i uint32
	for i = 0; i < im.Sites(); i++ {
		energy := im.Energy(i)
		sqrSum += energy * energy
	}

	avg := AverageEnergy(im)
	return (sqrSum / 4.0 - avg * avg) / float64(im.Sites())
}


func Chi(im IsingModel) float64 {
	sqrSum := 0.0

	var i uint32
	for i = 0; i < im.Sites(); i++ {
		mag := im.Magnetization(i)
		sqrSum += mag * mag
	}

	avg := AverageMag(im)
	return (sqrSum - avg * avg) / float64(im.Sites())
}
