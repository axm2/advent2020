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
	counter := 0
	lines := strings.Split(text, "\r\n\r\n")
	//magicFields := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, line := range lines {
		line = strings.Replace(line, "\r\n", " ", -1)
		h := make(map[string]bool)
		fields := strings.Split(line, " ")
		for _, field := range fields {
			fieldSplit := strings.Split(field, ":")
			h[fieldSplit[0]] = true
		}
		if h["byr"] && h["iyr"] && h["eyr"] && h["hgt"] && h["hcl"] && h["ecl"] && h["pid"] {
			counter++
		}
	}
	fmt.Println(counter)
}
