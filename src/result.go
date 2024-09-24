package mathskills

import "fmt"


// displayResults prints the calculated statistical data

func DisplayResults(average, median, variance, deviation float64) {
	fmt.Printf("Average: %d\n", Round(average))
	fmt.Printf("Median: %d\n", Round(median))
	fmt.Printf("Variance: %d\n", Round(variance))
	fmt.Printf("Standard Deviation: %d\n", Round(deviation))
}
