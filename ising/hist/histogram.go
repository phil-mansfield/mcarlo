package hist

import (
	"math"

	"bitbucket.org/phil-mansfield/num"
)

type Histogram struct {
	Hist []int
	Start, Width float64
}

func New(length int, start, width float64) *Histogram {
	return &Histogram{ make([]int, length), start, width }
}

func (h *Histogram) AddFloat(x float64) {
	i := int((x - h.Start) / h.Width)
	if i < 0 {
		h.Hist[0] += 1
	} else if i >= len(h.Hist) {
		h.Hist[len(h.Hist) - 1] += 1
	} else {
		h.Hist[i] += 1
	}
}

func (h *Histogram) Energy(t0, t float64) float64 {
	num, den := 0.0, 0.0
	b0, b := 1/t0, 1/t

	for i := 0; i < len(h.Hist); i++ {
		if h.Hist[i] == 0 { continue }

		E := h.Start + h.Width * float64(i)
		num += float64(h.Hist[i]) * E * math.Exp(E * (b0 - b))
		den += float64(h.Hist[i]) * math.Exp(E * (b0 - b))
	}

	return num / den
}

func (h *Histogram) SpecificHeat(t0, t float64) float64 {
	num, den := 0.0, 0.0
	b0, b := 1/t0, 1/t

	for i := 0; i < len(h.Hist); i++ {
		if h.Hist[i] == 0 { continue }

		E := h.Start + h.Width * float64(i)
		num += float64(h.Hist[i]) * E * E * math.Exp(E * (b0 - b))
		den += float64(h.Hist[i]) * math.Exp(E * (b0 - b))
	}
	
	E := h.Energy(t0, t) // This is non-optimal.

	return (num / den - E * E) /  (2.0 * t * t)
}

func (h *Histogram) CriticalTemp(t0, minTemp, maxTemp float64) float64 {
	specificHeat1D := func(temp float64) float64 {
		return h.SpecificHeat(t0, temp)
	}
	return num.Maximum(specificHeat1D, minTemp, maxTemp)
}