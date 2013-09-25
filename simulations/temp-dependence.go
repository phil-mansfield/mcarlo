package main

import (
	"fmt"
	"time"

	"bitbucket.org/phil-mansfield/table"
	"github.com/phil-mansfield/mcarlo/ising"
	"github.com/phil-mansfield/mcarlo/mcrand"
)

const (
	initialSweeps = 100
	maxTemp = 4.0
	minTemp = 1.0
	points = 51.0
	runs = 5
	trials = 500
	widthPow = 2
)

func main() {
	fmt.Println("Run statistics:")
	fmt.Printf("  %15s %5d\n", "Grid Width", 1<<widthPow)
	fmt.Printf("  %15s %5d\n", "Initial Sweeps", initialSweeps)
	fmt.Printf("  %15s %5d\n", "Trial Sweeps", trials)
	fmt.Printf("  %15s %5d\n", "Runs", runs)
	fmt.Printf("  %15s %5d\n", "Plot Points", int(points))

	t0 := float64(time.Now().UnixNano())

	infos := make([]*ising.SweepInfo, runs)
	for i := 0; i < runs; i++ {
		infos[i] = &ising.SweepInfo{
			ising.New(ising.Grid2DType, widthPow),
			ising.RandomSweep,
			mcrand.New(mcrand.Xorshift, int64(time.Now().Nanosecond())),
			minTemp,
		}
	}

	t := table.NewOutTable("Temp", "Energy", "Mag", "S. Heat", "Chi")
	stats := make([]*ising.SweepStats, runs)

	print("Running simulation")
	for point := 0.0; point < points; point++ {
		print(".")

		for i, info := range infos {
			info.Temp = minTemp + (maxTemp - minTemp) * point / points
			stats[i] = ising.SweepStatistics(info, initialSweeps, trials)
		}

		stat := ising.AvgSweepStats(stats)

		t.AddRow(infos[0].Temp, stat.E, stat.Mag, stat.C, stat.X)
	}
	println()

	t.Write(table.KeepHeader, fmt.Sprintf("temp%d.table", 1<<widthPow))

	t1 := float64(time.Now().UnixNano())
	fmt.Printf("Run time: %.3gs\n", (t1 - t0) / 1e9)
}