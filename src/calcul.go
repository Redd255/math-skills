package mathskills

import "math"

func Average(Sorted []float64) float64 {
	var sum float64
	var average float64
	for _, value := range Sorted {
		sum += value
	}
	average = sum / float64(len(Sorted))
	return average
}
func Mediane(Sorted []float64) float64 {
	var mediane float64

	if len(Sorted)%2 == 0 {
		mediane = (Sorted[(len(Sorted)/2-1)] + Sorted[(len(Sorted)/2)]) / 2
	} else {
		mediane = (Sorted[len(Sorted)/2])
	}
	return mediane
}

func Variance(Sorted []float64) (float64, float64) {
	var res float64
	var variance float64
	var deviation float64
	average := Average(Sorted)
	for _, value := range Sorted {
		res = (value - average) * (value - average)
		variance += res
	}
	variance = variance / float64(len(Sorted))
	deviation = math.Sqrt(variance)

	return variance, deviation
}
func Round(x float64) int {
	var rounded int
	if x >= 0 {
		rounded = int(x + 0.5)
	} else {
		rounded = int(x - 0.5)
	}
	return rounded
}
