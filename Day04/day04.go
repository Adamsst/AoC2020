package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const inputPath = "input.txt"

type Passport struct {
	Byr int
	Iyr int
	Eyr int
	Hgt string
	Hcl string
	Ecl string
	Pid string
	Cid int
}

func main() {
	var part2 = 0

	input, err := readLines(inputPath)
	if err != nil {
		fmt.Println("error reading input")
		return
	}

	var passports = make([]Passport, 1)

	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			if i == len(input)-1 {
				break
			}
			passports = append(passports, Passport{})
			continue
		} else {
			passports[len(passports)-1] = processStr(input[i], passports[len(passports)-1])
		}

	}
	for i := 0; i < len(passports); i++ {
		if checkIfValid(passports[i]) {
			part2++
		}
	}

	fmt.Printf("Part2: %d \n", part2)

}

// read lines of a file into a slice
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

func processStr(in string, p Passport) Passport {
	var splitInput = strings.Split(in, " ")

	for i := 0; i < len(splitInput); i++ {
		splitAgain := strings.Split(splitInput[i], ":")

		switch splitAgain[0] {
		case "byr":
			p.Byr = getInt(splitAgain[1])
		case "iyr":
			p.Iyr = getInt(splitAgain[1])
		case "eyr":
			p.Eyr = getInt(splitAgain[1])
		case "hgt":
			p.Hgt = splitAgain[1]
		case "hcl":
			p.Hcl = splitAgain[1]
		case "ecl":
			p.Ecl = splitAgain[1]
		case "pid":
			p.Pid = splitAgain[1]
		case "cid":
			p.Cid = getInt(splitAgain[1])
		default:
			fmt.Println("Default Case on passport parameter!")
		}

	}
	return p
}

func getInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Err converting str to int!")
	}
	return i
}

// This is making assumption that 0 value is not valid
func checkIfValid(p Passport) bool {
	if !intBetween(1920, 2002, p.Byr) {
		return false
	}
	if !intBetween(2010, 2020, p.Iyr) {
		return false
	}
	if !intBetween(2020, 2030, p.Eyr) {
		return false
	}
	if len(p.Pid) != 9 {
		return false
	}

	_, err := strconv.Atoi(p.Pid)
	if err != nil {
		return false
	}

	if !(p.Ecl == "amb" || p.Ecl == "blu" || p.Ecl == "brn" || p.Ecl == "gry" || p.Ecl == "grn" || p.Ecl == "hzl" || p.Ecl == "oth") {
		return false
	}
	if !checkHeight(p.Hgt) {
		return false
	}
	if !checkHair(p.Hcl) {
		return false
	}

	return true
}

func intBetween(low, high, in int) bool {
	if in < low {
		return false
	}
	if in > high {
		return false
	}
	return true
}

func checkHeight(height string) bool {
	if strings.HasSuffix(height, "cm") {
		intHeight, _ := strconv.Atoi(strings.ReplaceAll(height, "cm", ""))
		return intBetween(150, 193, intHeight)
	}
	if strings.HasSuffix(height, "in") {
		intHeight, _ := strconv.Atoi(strings.ReplaceAll(height, "in", ""))
		return intBetween(59, 76, intHeight)
	}
	return false
}

func checkHair(hair string) bool {
	if len(hair) != 7 {
		return false
	}
	if !strings.HasPrefix(hair, "#") {
		return false
	}

	runes := []rune(hair)

	matched, err := regexp.Match(`[a-f0-9]+`, []byte(string(runes[1:])))
	if err != nil {
		fmt.Println("Hair regexp err: " + err.Error())
		return false
	}
	return matched

}
