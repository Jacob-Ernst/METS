package conversions

import "math"

// For converting lb to kg.
const kgConversion float64 = 0.45359237

// Convert pound to kilograms and round the result.
func PoundsToMetric(lb float64) float64 {
	return roundWeight(lb * kgConversion)
}

// Rounds float64 to percision 2.
func roundWeight(weight float64) float64 {
	return math.Round(weight*100.0) / 100.0
}
