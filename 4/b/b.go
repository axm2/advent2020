package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
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
		h := make(map[string]string)
		fields := strings.Split(line, " ")
		for _, field := range fields {
			fieldSplit := strings.Split(field, ":")
			h[fieldSplit[0]] = fieldSplit[1]
		}
		byr, _ := strconv.Atoi(h["byr"])
		iyr, _ := strconv.Atoi(h["iyr"])
		eyr, _ := strconv.Atoi(h["eyr"])
		hgtNum := 0
		hgtUnit := ""
		if hm, _ := regexp.Match(`[0-9]{2,3}(cm|in)`, []byte(h["hgt"])); hm != false {
			hgtNum, _ = strconv.Atoi(h["hgt"][:len(h["hgt"])-2])
			hgtUnit = h["hgt"][len(h["hgt"])-2:]
		}
		hcl := h["hcl"]
		ecl := h["ecl"]
		pid := h["pid"]

		bbyr := byr >= 1920 && byr <= 2002
		biyr := iyr >= 2010 && iyr <= 2020
		beyr := eyr >= 2020 && eyr <= 2030
		bhgt := false
		if hgtUnit == "cm" && hgtNum >= 150 && hgtNum <= 193 {
			bhgt = true
		}
		if hgtUnit == "in" && hgtNum >= 59 && hgtNum <= 76 {
			bhgt = true
		}
		bhcl, _ := regexp.Match(`#([0-9]|[a-f]){6}`, []byte(hcl))
		becl := false
		if ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth" {
			becl = true
		}
		bpid, _ := regexp.Match(`^[0-9]{9}$`, []byte(pid))
		fmt.Println(line)
		fmt.Print(bbyr)
		fmt.Print(biyr)
		fmt.Print(beyr)
		fmt.Print(bhgt)
		fmt.Print(bhcl)
		fmt.Print(becl)
		fmt.Println(bpid)
		fmt.Println("-----------------------------------------------------------")
		if bbyr && biyr && beyr && bhgt && bhcl && becl && bpid {
			counter++
		}
	}
	fmt.Println(counter)
}
