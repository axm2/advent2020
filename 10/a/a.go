package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return lines, err
		}
		lines = append(lines, x)
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
	lines, err := readLines("../test.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	lines = append(lines, 0)
	sort.Ints(lines)
	threes := 1
	ones := 0
	for i := 0; i < len(lines)-1; i++ {
		gap := lines[i+1] - lines[i]
		if gap == 1 {
			ones++
		} else if gap == 3 {
			threes++
		}
	}
	fmt.Println(lines)
	fmt.Println(threes)
	fmt.Println(ones)
	fmt.Println(threes * ones)
}
