package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func removeDuplicateValues(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	lines, err := readLines("../input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	h := make(map[string][]string)
	baglines := []string{"bags.", "bags", "bag.", "bag"}
	for _, line := range lines {
		cleanedLine := line
		for _, bagline := range baglines {
			cleanedLine = strings.Replace(cleanedLine, bagline, "", -1)
		}
		splitted := strings.Split(cleanedLine, "contain")
		outerBag := strings.Replace(splitted[0], " ", "", -1)
		innerBags := strings.Split(splitted[1], ",")
		for _, innerbag := range innerBags {
			cleanedib := strings.Replace(innerbag, " ", "", -1)
			h[outerBag] = append(h[outerBag], cleanedib)
		}
	}
	fmt.Println(h)
	// bags := 0
	// for len(h["shinygold"]) > 0 {
	// 	pop := h["shinygold"][0]
	// 	bags = append(bags, pop)
	// 	h["shinygold"] = h["shinygold"][1:]
	// 	if len(h[pop]) > 0 {
	// 		h["shinygold"] = append(h["shinygold"], h[pop]...)
	// 	}
	// }
	// fmt.Println(len(removeDuplicateValues(bags)))
}
