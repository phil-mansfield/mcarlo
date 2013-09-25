package main

import (
	"fmt"
	"math"
	"time"

	"bitbucket.org/phil-mansfield/table"

	"github.com/phil-mansfield/mcarlo/ising"
	"github.com/phil-mansfield/mcarlo/mcrand"
)

const (
	initialSweeps = 100
	minSampleTemp = 2.0
	minTemp = 1.0
	maxTemp = 7.0
	tempStep = 0.01
	minTrialsPow = 1
	maxTrialsPow = 6
	minPow uint32 = 1
	maxPow uint32 = 6
)

// This is too monolithic and should be pruned.
func main() {
	sampleTemp := minSampleTemp

	maximumTable := table.NewOutTable("L", "1/L", 
		"Critical Temperature", "Max Specific Heat", "Max Specific Heat Per Site")
	maximumTable.SetHeader(fmt.Sprintf("Created with %g trials per grid.",
		math.Pow(10, maxTrialsPow)))

	for widthPow := minPow; widthPow <= maxPow; widthPow++ {
		cols := make([]string, 1 + 2 * (maxTrialsPow - minTrialsPow + 1))
		cols[0] = "Temperature"

		fcols := make([][]float64, int((maxTemp - minTemp) / tempStep))
		for i := 0; i < len(fcols); i++ {
			fcols[i] = make([]float64, 1 + 2 * (maxTrialsPow - minTrialsPow + 1))
		}

		secs1 := 0.0
		secs2 := 0.0

		fmt.Printf("\n%d-Width Grid:\n", 1<<widthPow)
		for i := 0; i <= maxTrialsPow - minTrialsPow; i++ {
			trials := int(math.Pow(10, float64(minTrialsPow + i)))
			fmt.Printf("  Running %d trial suite\n", trials)

			cols[2 * i + 1] = fmt.Sprintf("E (%d)", trials)
			cols[2 * i + 2] = fmt.Sprintf("C (%d)", trials)

			info := &ising.SweepInfo{
				ising.New(ising.Grid2DType, uint32(widthPow)),
				ising.IterSweep,
				mcrand.New(mcrand.Xorshift, int64(time.Now().UnixNano())),
				sampleTemp,
			}
			
			t0 := time.Now().UnixNano()
			ising.SweepStatistics(info, initialSweeps, trials)
			t1 := time.Now().UnixNano()
			secs1 += float64(t1 - t0) / 1e9

			h := info.Im.EnergyHistogram()

			minValidTemp, maxValidTemp := maxTemp, minTemp
			for j := 0; j < len(fcols); j++ {
				temp := minTemp + float64(j) * tempStep
		
				E := h.Energy(sampleTemp, temp)
				C := h.SpecificHeat(sampleTemp, temp)

				if !math.IsNaN(C) && !math.IsInf(C, 0) {
					if temp > maxValidTemp {
						maxValidTemp = temp
					}
					if temp < minValidTemp {
						minValidTemp = temp
					}
				}

				fcols[j][2 * i + 1] = E / float64(info.Im.Sites())
				fcols[j][2 * i + 2] = C / float64(info.Im.Sites())
			}

			if i + minTrialsPow == maxTrialsPow {
				tCrit := h.CriticalTemp(sampleTemp, minValidTemp, maxValidTemp)
				cCrit := h.SpecificHeat(sampleTemp, tCrit)
				maximumTable.AddRow(float64(int(1 << widthPow)),
					1.0 / float64(int(1 << widthPow)), tCrit,
					cCrit, cCrit / float64(info.Im.Sites()))

				sampleTemp = tCrit
			}
		
			t2 := time.Now().UnixNano()
			secs2 += float64(t2 - t1) / 1e9
		}

		headerString := fmt.Sprintf(
			"%16s: %g\n%16s: %d\n%16s: %d\n" + 
				"%16s: %s\n%16s: %s\n%16s: %g\n%20s: %g",
			"Sample Temperature", sampleTemp,
			"Initial Sweeps", initialSweeps,
			"Width", 1<<widthPow,
			"Generator", "Xorshift",
			"Sweep Pattern", "IterSweep",
			"Total Sweep Time", secs1,
			"Total Histogram Time", secs2)
		
		t := table.NewOutTable(cols...)

		t.SetHeader(headerString)

		for i := 0; i < len(fcols); i++ {
			if !math.IsNaN(fcols[i][1]) {
				fcols[i][0] = minTemp + tempStep * float64(i)
				t.AddRow(fcols[i]...)
			}
		}

		t.Write(table.KeepHeader, 
			fmt.Sprintf("tables/iter/hist-temp%d.table", 1<<widthPow))
	}

	maximumTable.Write(table.KeepHeader, "tables/iter/hist-maximum.table")
}