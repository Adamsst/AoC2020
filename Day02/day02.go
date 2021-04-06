package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputPath = "input.txt"

type password struct {
	letter string
	min    int
	max    int
	pass   string
}

func main() {
	var part1 int
	var part2 int

	input, err := readLines(inputPath)
	if err != nil {
		fmt.Println("error reading input")
		return
	}

	var inputPasswords = make([]password, len(input))
	part1 = len(input)
	part2 = len(input)

	for i, p := range input {
		inputPasswords[i] = makePassword(p)
		if strings.Count(inputPasswords[i].pass, inputPasswords[i].letter) < inputPasswords[i].min {
			part1--
			continue
		}
		if strings.Count(inputPasswords[i].pass, inputPasswords[i].letter) > inputPasswords[i].max {
			part1--
			continue
		}
	}

	// part 2
	for _, p := range inputPasswords {
		if string(p.pass[p.min-1]) == p.letter && string(p.pass[p.max-1]) == p.letter {
			part2--
			continue
		}
		if string(p.pass[p.min-1]) != p.letter && string(p.pass[p.max-1]) != p.letter {
			part2--
			continue
		}
	}

	fmt.Printf("Part1: %d \n", part1)
	fmt.Printf("Part2: %d \n", part2)
}

func makePassword(input string) password {
	spaced := strings.Split(input, " ")
	minMax := strings.Split(spaced[0], "-")
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])

	return password{
		letter: string(spaced[1][0]),
		min:    min,
		max:    max,
		pass:   spaced[2],
	}
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
