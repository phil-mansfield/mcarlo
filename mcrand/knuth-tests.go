package mcrand

import (
	rec "github.com/phil-mansfield/gorecipies"
)

func frequencyTest(gen *Generator, d, iters int64) float64 {
	bins := make([]int64, d)
	for n := int64(0); n < iters; n++ {
		bins[gen.UniformInt(0, d - 1)]++
	}

	probs := make([]float64, d)
	for i, _ := range probs {
		probs[i] = 1 / float64(d)
	}
	_, prob := rec.ChiSqr(bins, probs)
	return prob
}

func FrequencyTest(gen *Generator, d, iters, chiNum int64) []float64 {
	chis := make([]float64, chiNum)

	for i := int64(0); i < chiNum; i++ {
		chis[i] = frequencyTest(gen, d, iters / chiNum)
	}

	return chis
}


func serialTest(gen *Generator, d, iters int64) float64 {
	bins := make([]int64, d * d)
	for n := int64(0); n < iters; n++ {
		x, y := gen.UniformInt(0, d - 1), gen.UniformInt(0, d - 1)
		bins[y * d + x]++
	}

	probs := make([]float64, d * d)
	for i, _ := range probs {
		probs[i] = 1 / float64(d * d)
	}

	_, prob := rec.ChiSqr(bins, probs)
	return prob
}

func SerialTest(gen *Generator, d, iters, chiNum int64) []float64 {
	chis := make([]float64, chiNum)

	for i := int64(0); i < chiNum; i++ {
		chis[i] = serialTest(gen, d, iters / chiNum)
	}

	return chis
}
