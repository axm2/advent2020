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

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func model(lines []string) []string {
	var linescpy []string
	for i := 0; i < len(lines); i++ {
		outer := ""
		for j := 0; j < len(lines[i]); j++ {
			counter := 0
			if i > 0 && lines[i-1][j] == '#' {
				// up
				counter++
			}
			if j > 0 && lines[i][j-1] == '#' {
				// left
				counter++
			}
			if j < len(lines[i])-1 && lines[i][j+1] == '#' {
				//right
				counter++
			}
			if i < len(lines)-1 && lines[i+1][j] == '#' {
				//down
				counter++
			}
			if i > 0 && j > 0 && lines[i-1][j-1] == '#' {
				//upleft
				counter++
			}
			if i > 0 && j < len(lines[i])-1 && lines[i-1][j+1] == '#' {
				//upright
				counter++
			}
			if i < len(lines)-1 && j > 0 && lines[i+1][j-1] == '#' {
				//downleft
				counter++
			}
			if j < len(lines[i])-1 && i < len(lines)-1 && lines[i+1][j+1] == '#' {
				//downright
				counter++
			}
			if lines[i][j] == 'L' && counter == 0 {
				//linescpy[i][j] = '#'
				outer += "#"
			} else if lines[i][j] == '#' && counter >= 4 {
				//linescpy[i][j] = 'L'
				outer += "L"
			} else {
				outer += string(lines[i][j])
			}
		}
		linescpy = append(linescpy, outer)
	}
	return linescpy
}

func main() {
	lines, err := readLines("../input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	before := lines
	after := model(lines)
	for !Equal(before, after) {
		before = model(before)
		after = model(before)
	}
	counter := 0
	for _, line := range after {
		for _, rune := range line {
			if rune == '#' {
				counter++
			}
		}
	}
	fmt.Println(counter)
}
