package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	result := 0
	for _, line := range lines {
		split := strings.Fields(line)
		// split[0] is our numbers
		// split[1][0] is the character to check for
		magicRune := rune(split[1][0])
		// split[2] is the password given

		nums := strings.Split(split[0], "-")
		// nums[0] and nums[1] are our numbers
		pos1, _ := strconv.Atoi(nums[0])
		pos2, _ := strconv.Atoi(nums[1])

		if (rune(split[2][pos1-1]) == magicRune) && (rune(split[2][pos2-1]) != magicRune) || (rune(split[2][pos1-1]) != magicRune) && (rune(split[2][pos2-1]) == magicRune) {
			result++
		}

		fmt.Println(line)
		fmt.Println(pos1)
		fmt.Println(pos2)
		fmt.Println(result)
		fmt.Println("--------------------")
	}
	fmt.Println(result)
}
