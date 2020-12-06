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
		line = strings.Replace(line, "\r\n", "", -1)
		counter := 0
		h := make(map[rune]bool)
		for _, rune := range line {
			if !h[rune] {
				h[rune] = true
				counter++
			}
		}
		sum = counter + sum
	}
	fmt.Println(sum)
}
