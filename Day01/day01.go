package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inputPath = "input.txt"
const goal = 2020

func main() {
	var part1 int
	var part2 int

	input, err := readLines(inputPath)
	if err != nil {
		fmt.Println("error reading input")
		return
	}

	var numInput = make([]int, len(input))
	for i, num := range input {
		// all ints in the file
		numInput[i], _ = strconv.Atoi(num)
	}

	for _, num1 := range numInput {
		for _, num2 := range numInput {
			if num1+num2 == goal {
				part1 = num1 * num2
			}

			if num1+num2 >= goal {
				continue
			}

			for _, num3 := range numInput {
				if num1+num2+num3 == goal {
					part2 = num1 * num2 * num3
				}
			}
		}
	}

	fmt.Printf("Part1: %d \n", part1)
	fmt.Printf("Part2: %d \n", part2)
	fmt.Println("Hello, World!")
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
