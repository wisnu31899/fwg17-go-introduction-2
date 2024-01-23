package main

import (
	"math"
)

func pembulatan(nilai float64) float64 {
	hasil := math.Round(nilai*10) / 10
	return hasil
}
