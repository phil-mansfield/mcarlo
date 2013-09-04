package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"bitbucket.org/phil-mansfield/table"
	"github.com/phil-mansfield/graph"
	"github.com/phil-mansfield/mcarlo/potts"
)

const (
	steps = 50
)

func sec(t0, t1 time.Time) float64 {
	return 1e-9 * float64(t1.UnixNano() - t0.UnixNano())
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <grid width> <out-file>\n")
		os.Exit(1)
	}
	
	width, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error with first argument: %s", err.Error())
		os.Exit(1)
	}

	timeTable := table.NewOutTable("Generate Edges", "Create Graph", "Union")
	dataTable := table.NewOutTable("Edge Frequency", "Largest Group Frac")
	dataTable.SetHeader(fmt.Sprintf(
		"Grid has continuous boundary conditions and %d x %d nodes.",
		width, width))

	for step := 0; step <= steps; step++ {
		prob := float64(step) / float64(steps)

		t0 := time.Now()
		us, vs := potts.GenerateEdges(potts.Grid2D, prob, int(width))

		t1 := time.Now()
		g := graph.New(uint32(width * width), us, vs)

		t2 := time.Now()
		g.Union()

		t3 := time.Now()

		largestGroup := g.Query(graph.Size, g.LargestGroup())
		dataTable.AddRow(prob, float64(largestGroup) / float64(width * width))
		timeTable.AddRow(sec(t0, t1), sec(t1, t2), sec(t2, t3))
	}

	timeTable.Print(table.KeepHeader)
	if err := dataTable.Write(table.KeepHeader, os.Args[2]); err != nil {
		fmt.Fprintf(os.Stderr, "I/O Error: %s\n", err.Error())
		os.Exit(1)
	}
}
