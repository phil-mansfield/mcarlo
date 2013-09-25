package main

import (
	"time"

	"bitbucket.org/phil-mansfield/table"
	"github.com/phil-mansfield/mcarlo/ising"
	"github.com/phil-mansfield/mcarlo/mcrand"
)


const (
	initialSweeps = 100
	temp = 2.7
	trials = 10 * 1000
	widthPow = 3
)

func main() {
	t := table.NewOutTable("Time Step", "C Correlation", "X Correlation")

	info := &ising.SweepInfo{
		ising.New(ising.Grid2DType, widthPow),
		ising.RandomSweep,
		mcrand.New(mcrand.GoCrypto, int64(time.Now().UnixNano())),
		temp,
	}

	stats := ising.SweepStatistics(info, initialSweeps, trials)

	fCs, _ := ising.Tau(stats.EHistory)
	fMags, _ := ising.Tau(stats.MagHistory)

	if len(fCs) < len(fMags) {
		for ; 0 < len(fMags) - len(fCs); {
			fCs = append(fCs, 0.0)
		}
	} else {
		for ; 0 < len(fCs) - len(fMags); {
			fMags = append(fMags, 0.0)
		}
	}

	info.Im.Print()

	println("Total steps:", len(fMags))

	for i := 0; i < len(fMags); i++ {
		t.AddRow(float64(i), fCs[i], fMags[i])
	}

	t.Write(table.KeepHeader, "tau.table")
}