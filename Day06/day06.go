package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const inputPath = "input.txt"

func main() {
	var p1Total = 0
	var p2Total = 0
	var newGroup = true
	var groupStr string

	input, err := readLines(inputPath)
	if err != nil {
		fmt.Println("Input read err: " + err.Error())
		return
	}

	for _, in := range input {
		if in == "" {
			p1Total += len(groupStr) // Account for group before starting a new group
			newGroup = true
			continue
		}
		if newGroup {
			groupStr = ""
			newGroup = false
		}
		groupStr = AddUniqueLetters(groupStr, in)
	}
	p1Total += len(groupStr) // Final group accounted for
	fmt.Println("P1:", p1Total)

	newGroup = true
	for _, in := range input {
		if in == "" {
			p2Total += len(groupStr) // Account for group before starting a new group
			newGroup = true
			continue
		}
		if newGroup {
			groupStr = in
			newGroup = false
		} else {
			groupStr = RemoveUniqueLetters(groupStr, in)
		}
	}
	p2Total += len(groupStr) // Final group accounted for
	fmt.Println("P2:", p2Total)
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

// AddUniqueLetters will add only unique letters from the input string to the given current string
func AddUniqueLetters(current, in string) string {
	for i := 0; i < len(in); i++ {
		if !strings.Contains(current, string(in[i])) {
			current += string(in[i])
		}
	}
	return current
}

// RemoveUniqueLetters will remove only unique letters from the current string if not in the input string
func RemoveUniqueLetters(current, in string) string {
	var charsToRemove string

	for i := 0; i < len(current); i++ {
		if !strings.Contains(in, string(current[i])) {
			charsToRemove += string(current[i])
		}
	}

	for i := 0; i < len(charsToRemove); i++ {
		current = strings.ReplaceAll(current, string(charsToRemove[i]), "")
	}

	return current
}
