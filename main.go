package main

import (
	"log"
	mathskills "mathskills/src"
)

func main() {
	// Read the data from the file
	data, err := mathskills.Read()
	if err != nil {
		log.Fatalf("Failed to read data: %v", err)
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
