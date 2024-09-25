package main

import (
	"fmt"
	mathskills "mathskills/src"
)

func main() {
	// Read the data from the file
	data, err := mathskills.Read()
	if err != nil {
		fmt.Println("Failed to read data:", err)
		return
	}

	// Sort the data
	sorted := mathskills.QuickSort(data)

	// Perform statistical calculations
	average := mathskills.Average(sorted)
	median := mathskills.Mediane(sorted)
	variance, deviation := mathskills.Variance(sorted)

	// Display results
	mathskills.DisplayResults(average, median, variance, deviation)
}
