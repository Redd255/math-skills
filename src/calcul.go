package mathskills

import "math"

func Average(Sorted []float64) float64 {
	if len(Sorted) == 0 {
		return 0
	}

	sum := 0.0
	for _, value := range Sorted {
		sum += value
	}

	return sum / float64(len(Sorted))
}

func Mediane(Sorted []float64) float64 {
	n := len(Sorted)
	if n == 0 {
		return 0
	}

	if n%2 == 0 {
		return (Sorted[n/2-1] + Sorted[n/2]) / 2
	}
	return Sorted[n/2]
}

func Variance(Sorted []float64) (float64, float64) {
	n := len(Sorted)
	if n == 0 {
		return 0, 0
	}

	average := Average(Sorted)
	var variance float64

	for _, value := range Sorted {
		diff := value - average
		variance += diff * diff
	}

	variance /= float64(n)
	deviation := math.Sqrt(variance)

	return variance, deviation
}

func Round(x float64) int {
	if x >= 0 {
		return int(x + 0.5)
	}
	return int(x - 0.5)
}
