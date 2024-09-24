package mathskils

func QuickSort(Population []float64) []float64 {
	var before, after, pivotlist []float64
	if len(Population) < 2 {
		return Population
	}
	first := Population[0]
	middle := Population[(len(Population)-1)/2]
	last := Population[len(Population)-1]
	var Pivot float64
	if (first > middle) && (first < last) {
		Pivot = first
	} else if (middle > first) && (middle < last) {
		Pivot = middle
	} else {
		Pivot = last
	}
	for _, value := range Population {
		if value < Pivot {
			before = append(before, value)
		} else if value > Pivot {
			after = append(after, value)
		} else {
			pivotlist = append(pivotlist, value)
		}
	}
	var Sorted []float64
	Sorted = append(Sorted, QuickSort(before)...)
	Sorted = append(Sorted, pivotlist...)
	Sorted = append(Sorted, QuickSort(after)...)
	return Sorted
}
