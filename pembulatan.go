package main

import (
	"math"
)

func pembulatan(nilai float64) (float64, float64) {
	rounded := math.Round(nilai*10) / 10
	remainder := rounded - math.Floor(rounded)
	return rounded, remainder
}
