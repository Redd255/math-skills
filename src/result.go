package mathskills

import "fmt"

// displayResults prints the calculated data

func DisplayResults(average, median, variance, deviation float64) {
	fmt.Println("Average:", Round(average))
	fmt.Println("Median:", Round(median))
	fmt.Println("Variance:", Round(variance))
	fmt.Println("Standard Deviation:", Round(deviation))
}
