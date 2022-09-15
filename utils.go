package main

import (
	"math"
	"path/filepath"
)

func fileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func degrees(radians float64) float64 {
	return radians * 180 / math.Pi
}

func radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
