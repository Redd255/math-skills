package mathskills

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Read() ([]float64, error) {
	// Read the content of the file
	content, err := os.ReadFile("data.txt")
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// Trim spaces and split the content into lines
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	var population []float64

	// Parse each line into a float64 value 
	for _, line := range lines {
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Printf("Skipping invalid value: %s (error: %v)", line, err)
			continue
		}
		population = append(population, value)
	}

	// Return the parsed population data
	return population, nil
}
