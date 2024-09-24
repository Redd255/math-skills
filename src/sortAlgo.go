package mathskills

func QuickSort(arr []float64) []float64 {
	if len(arr) < 2 {
		return arr
	}

	// Choose pivot as the middle element
	pivot := arr[len(arr)/2]
	var less, equal, greater []float64

	// Partitioning the array based on pivot
	for _, v := range arr {
		switch {
		case v < pivot:
			less = append(less, v)
		case v > pivot:
			greater = append(greater, v)
		default:
			equal = append(equal, v)
		}
	}

	// Recursively sort less and greater arrays, then combine
	return append(append(QuickSort(less), equal...), QuickSort(greater)...)
}
