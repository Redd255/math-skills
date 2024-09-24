package mathskils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Read() []float64 {
	var Population []float64

	content, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal("couldn't read file")
	}

	Split := strings.Split(strings.TrimSpace(string(content)), "\n")

	for _, v := range Split {
		s, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Print("Parse failed", string(v), "in this value")
		} else {
			Population = append(Population, s)
		}
	}
	return Population
}
