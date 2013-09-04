package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/phil-mansfield/mcarlo/potts"
)

const (
	gridWidth = 4
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Proper usage: %s <probability>\n")
		os.Exit(1)
	}

	prob, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	} else if prob > 1 {
		prob = 1
	} else if prob < 0 {
		prob = 0
	}

	us, vs := potts.GenerateEdges(potts.Grid2D, prob, gridWidth)

	fmt.Printf("For a grid of width %d with %d nodes the following edges " +
		"were generated:\n", gridWidth, gridWidth * gridWidth)

	for i := 0; i < len(us); i++ {
		fmt.Printf("  (%d, %d)\n", us[i], vs[i])
	}

	if len(us) == 0 {
		fmt.Println("nil")
	}
}
