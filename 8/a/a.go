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
	lines, err := readLines("../input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	i := 0
	accumulator := 0
	h := make(map[int]bool)
	for !h[i] {
		if string(lines[i][:3]) == "nop" {
			h[i] = true
			i++
		} else if string(lines[i][:3]) == "acc" {
			h[i] = true
			add, _ := strconv.Atoi(lines[i][4:])
			accumulator += add
			i++
		} else if string(lines[i][:3]) == "jmp" {
			h[i] = true
			add, _ := strconv.Atoi(lines[i][4:])
			i += add
		}
	}
	fmt.Println(accumulator)
}
