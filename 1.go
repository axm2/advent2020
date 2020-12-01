package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	lines, err := readLines("1.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	magicNumber := 2020
	h := make(map[int]bool)
	// part A
	/*
		for _, firstline := range lines {
			firstnum, _ := strconv.Atoi(firstline)
			h[firstnum] = true
			secondnum := magicNumber - firstnum
			if h[secondnum] {
				fmt.Println(firstnum * secondnum)
				return
			}

		}
	*/
	for _, firstline := range lines {
		firstnum, _ := strconv.Atoi(firstline)
		h[firstnum] = true
		partialsum := magicNumber - firstnum
		// partialsum becomes our new magicNumber
		for _, secondline := range lines {
			secondnum, _ := strconv.Atoi(secondline)
			thirdnum := partialsum - secondnum
			if h[thirdnum] {
				fmt.Println(firstnum * secondnum * thirdnum)
				return
			}
		}
	}
}
