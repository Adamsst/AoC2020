package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const rowChars = 7
const colChars = 3
const inputPath = "input.txt"

func main() {
	var rowSlice = make([]int, 128)
	var colSlice = make([]int, 8)
	var part2 = make([]int, 0)
	var highID = 0

	for i := 0; i < 128; i++ {
		rowSlice[i] = i
	}
	for i := 0; i < 8; i++ {
		colSlice[i] = i
	}

	input, err := readLines(inputPath)
	if err != nil {
		fmt.Println("Input read err: " + err.Error())
		return
	}

	for _, in := range input {
		temp := analyzeRow(in, rowSlice, colSlice)
		part2 = append(part2, temp)
		if temp > highID {
			highID = temp
		}
	}

	fmt.Printf("P1: %d \n", highID)
	sort.Ints(part2)
	// This should print the missing seat twice
	for i := 1; i < len(part2)-1; i++ {
		if part2[i-1] != part2[i]-1 {
			fmt.Println(part2[i] - 1)
		}
		if part2[i+1] != part2[i]+1 {
			fmt.Println(part2[i] + 1)
		}
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

func analyzeRow(input string, rowSlice []int, colSlice []int) int {
	for i := 0; i < rowChars; i++ {
		if input[i] == 'F' {
			rowSlice = rowSlice[:len(rowSlice)/2]
		} else {
			rowSlice = rowSlice[len(rowSlice)/2:]
		}
	}

	for i := rowChars; i < rowChars+colChars; i++ {
		if input[i] == 'L' {
			colSlice = colSlice[:len(colSlice)/2]
		} else {
			colSlice = colSlice[len(colSlice)/2:]
		}
	}

	return (rowSlice[0] * 8) + colSlice[0]
}
