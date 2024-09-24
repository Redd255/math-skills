package main

import (
	"fmt"
	"io"
	"log"
	mathskils "mathskills/src"
	"os"
	"strconv"
)

func main() {
	Data := mathskils.Read()
	Sorted := mathskils.QuickSort(Data)
	resfile, err := os.Create("result/"+"result.txt")
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
	average := mathskils.Average(Sorted)
	median := mathskils.Mediane(Sorted)
	variance, deviation := mathskils.Variance(Sorted)
	fmt.Println("Average:", mathskils.Round(average))
	fmt.Println("Median:", mathskils.Round(median))
	fmt.Println("Variance:", mathskils.Round(variance))
	fmt.Println("Standard Deviation:", mathskils.Round(deviation))
}
