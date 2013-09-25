package main

import (
	"math"
	"time"

	"bitbucket.org/phil-mansfield/table"
	"github.com/phil-mansfield/mcarlo/ising"
)

func sec(t0, t1 time.Time) float64 {
    return 1e-9 * float64(t1.UnixNano() - t0.UnixNano())
}

func avg(xs []float64) float64 {
	sum := 0.0
	for _, x := range xs {
		sum += x
	}
	return sum / float64(len(xs))
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

func main() {
	// Temperature-dependent curves:

	t := table.NewOutTable("Temp", "Energy", "Magnetization",  "C", "Chi")
	im := ising.New(ising.Grid2DType, 16)

	t0 := time.Now()
	for temp := 0.0; temp < 5.0; temp += 0.1 {
		ising.Equilibriate(im, temp, 100)

		es := make([]float64, 100)
		ms := make([]float64, 100)

		for i := 0; i < len(es); i++ {
			ising.Equilibriate(im, temp, 1000 / len(es))
			es[i] = ising.AverageEnergy(im)
			ms[i] = math.Abs(ising.AverageMag(im))
		}
		t.AddRow(temp, avg(es), avg(ms),
			16 * 16 * variance(es) / temp / temp,
			16 * 16 * variance(ms) / temp)
	}
	t1 := time.Now()

	t.Write(table.KeepHeader, "temp.table")
	println("ising seconds:", sec(t0, t1))

	// Hysteresis loop:

	t = table.NewOutTable("h", "Magnetization")
	temp := 2.15

	t2 := time.Now()

	im = ising.New(ising.Grid2DType, 32)

	update := func(mag float64) {
		im.SetMag(mag)
		ising.Equilibriate(im, temp, 20)
		t.AddRow(mag, ising.AverageMag(im), ising.AverageEnergy(im),
			ising.SpecificHeat(im), ising.Chi(im))
	}

	for mag := -1.5; mag <= 1.5; mag += 0.03 { update(mag) }
	for mag := 1.5; mag >= -1.5; mag -= 0.03 { update(mag) }

	t3 := time.Now()
	println("hysteresis seconds:", sec(t2, t3))

	t.Write(table.KeepHeader, "mag.table")
}
