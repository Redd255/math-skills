package main

import (
	"fmt"
	"io"
	"log"
	mathskills "mathskills/src"
	"os"
	"strconv"
)

func main() {
	//read the file
	Data, err := mathskills.Read()
	if err != nil {
		log.Fatal(err)
	}

	//sort the data
	Sorted := mathskills.QuickSort(Data)
	resfile, err := os.Create("result/" + "result.txt")
	if err != nil {
		log.Fatal("Error creating file : ", err)
	}
	defer resfile.Close()
	var s string
	for _, v := range Sorted {
		s += strconv.FormatFloat(v, 'f', 0, 64) + " "
	}
	_, err = io.WriteString(resfile, s)
	if err != nil {
		log.Fatal("Error writing to file : ", err)
	}
	average := mathskills.Average(Sorted)
	median := mathskills.Mediane(Sorted)
	variance, deviation := mathskills.Variance(Sorted)
	fmt.Println("Average:", mathskills.Round(average))
	fmt.Println("Median:", mathskills.Round(median))
	fmt.Println("Variance:", mathskills.Round(variance))
	fmt.Println("Standard Deviation:", mathskills.Round(deviation))
}
