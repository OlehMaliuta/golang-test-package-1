package test1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func CreateFileWithSquareRoots(filename string, numbers []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, num := range numbers {
		root := math.Sqrt(float64(num))
		_, err := writer.WriteString(fmt.Sprintf("%.5f\n", root))
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func ReadSquareRootsFromFile(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var roots []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		root, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
		if err != nil {
			return nil, err
		}
		roots = append(roots, root)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return roots, nil
}

func ProductOfComponents(numbers []float64) float64 {
	product := 1.0
	for _, num := range numbers {
		product *= num
	}
	return product
}
