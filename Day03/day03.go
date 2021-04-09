package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const inputPath = "input.txt"

func main() {
	var part1 = 0
	var part2 int

	input, err := readLines(inputPath)
	if err != nil {
		fmt.Println("error reading input")
		return
	}

	part1 = calcTrees(input, 3, 1)

	part2 = part1
	part2 *= calcTrees(input, 1, 1)
	part2 *= calcTrees(input, 5, 1)
	part2 *= calcTrees(input, 7, 1)
	part2 *= calcTrees(input, 1, 2)

	fmt.Printf("Part1: %d \n", part1)
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

// calculate the number of trees "#" and use x and y movement as slope to scale input
func calcTrees(input []string, xMovement, yMovement int) int {
	var treeCount = 0
	var height = len(input)
	var width = len(input[0])
	var repetitions = math.Ceil(float64(height)/float64(width)) * math.Ceil(float64(xMovement)/float64(yMovement))

	for i := 0; i < len(input); i++ {
		var baseString = input[i]
		for j := 0; j < int(repetitions); j++ {
			input[i] = input[i] + baseString
		}
	}

	var x = 0
	var y = 0
	for y < len(input) {
		if string(input[y][x]) == "#" {
			treeCount++
		}

		x += xMovement
		y += yMovement
	}
	return treeCount
}
