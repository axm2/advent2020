package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func main() {
	lines, err := readLines("../input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	// find max
	max := 0.0
	min := math.MaxFloat64
	occupiedSeats := make(map[int]bool)
	for _, line := range lines {
		lowRow := 0
		highRow := 127
		lowCol := 0
		highCol := 7
		for _, rune := range line {
			if rune == 'F' {
				highRow = ((highRow - lowRow + 1) / 2) + lowRow - 1
			}
			if rune == 'B' {
				lowRow = (highRow+1-lowRow)/2 + lowRow
			}
			if rune == 'L' {
				highCol = ((highCol - lowCol + 1) / 2) + lowCol - 1
			}
			if rune == 'R' {
				lowCol = (highCol+1-lowCol)/2 + lowCol
			}
		}
		ID := lowRow*8 + lowCol
		occupiedSeats[ID] = true
		max = math.Max(float64(ID), max)
		min = math.Min(float64(ID), min)
	}
	for i := min; i <= max; i++ {
		if !occupiedSeats[int(i)] {
			fmt.Println(i)
			return
		}
	}

}
