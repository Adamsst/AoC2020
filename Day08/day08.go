package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputPath = "input.txt"

func main() {
	var acc = 0
	var index = 0
	var noLoop = true

	input, err := readLines(inputPath)
	if err != nil {
		fmt.Println("Input read err: " + err.Error())
		return
	}

	var visits = make(map[int]bool, len(input))

	for noLoop {
		if _, exists := visits[index]; exists {
			fmt.Printf("P1: %d\n", acc)
			noLoop = false
			break
		}
		visits[index] = true
		acc, index, _ = instruct(input, acc, index)
	}

	var maybeDefects = make([]int, 0)
	//Gather potential defective instructions
	for i := 0; i < len(input); i++ {
		if input[i][0] == 'n' || input[i][0] == 'j' {
			maybeDefects = append(maybeDefects, i)
		}
	}

	for i := 0; i < len(maybeDefects); i++ {
		index = 0
		noLoop = true
		acc = 0
		temp := make([]string, len(input))
		var end = false
		var visits2 = make(map[int]bool, len(input))
		copy(temp, input)

		//Replace the instruction at the potential defect index
		if temp[maybeDefects[i]][0] == 'j' {
			temp[maybeDefects[i]] = "nop" + temp[maybeDefects[i]][3:]
		} else if temp[maybeDefects[i]][0] == 'n' {
			temp[maybeDefects[i]] = "jmp" + temp[maybeDefects[i]][3:]
		}

		for noLoop {
			if _, exists := visits2[index]; exists {
				noLoop = false
				break
			}
			visits2[index] = true
			acc, index, end = instruct(temp, acc, index)
			if end {
				break
			}
		}

		if end {
			break
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

func instruct(instructions []string, acc, index int) (int, int, bool) {
	var parts = strings.Split(instructions[index], " ")
	var end = false
	numPart, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("instruct err converting str to int! " + err.Error())
	}

	switch parts[0] {
	case "nop":
		index++
	case "acc":
		acc += numPart
		index++
	case "jmp":
		index += numPart
	}

	if index > len(instructions) {
		index = index % len(instructions)
	} else if index < 0 {
		index = len(instructions) - (index % len(instructions))
	}

	if index == len(instructions) {
		end = true
		fmt.Printf("P2: %d\n", acc)
	}

	return acc, index, end
}
