package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	// windows does \r and \n
	lines := strings.Split(text, "\r\n\r\n")
	//magicFields := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	sum := 0
	for _, line := range lines {
		line = strings.Replace(line, "\r\n", " ", -1)
		groupAns := strings.Split(line, " ")
		h := make(map[rune]int)
		groupSize := len(groupAns)
		counter := 0
		for _, indivAns := range groupAns {
			for _, rune := range indivAns {
				if h[rune] == groupSize-1 {
					counter++
				}
				h[rune]++
			}
		}
		sum = counter + sum
	}
	fmt.Println(sum)
}
