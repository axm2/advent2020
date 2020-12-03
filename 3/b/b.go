package main

import (
	"bufio"
	"fmt"
	"log"
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

	idxx := [5]int{1, 3, 5, 7, 1}
	idxy := [5]int{1, 1, 1, 1, 2}
	xRight := [5]int{1, 3, 5, 7, 1}
	yDown := [5]int{1, 1, 1, 1, 2}
	slopeResults := [5]int{0, 0, 0, 0, 0}

	for i := 0; i < len(idxx); i++ {
		for j := 0; j < len(lines); j++ {
			if j == idxy[i] {
				// if this is a line we should consider...
				if lines[j][idxx[i]%len(lines[j])] == '#' {
					slopeResults[i]++
				}
				idxx[i] = idxx[i] + xRight[i]
				idxy[i] = idxy[i] + yDown[i]
			}
		}
	}

	product := 1
	for _, num := range slopeResults {
		product *= num
	}
	fmt.Println(product)
}
