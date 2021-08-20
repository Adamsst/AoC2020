package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inputPath = "input.txt"
const preamble = 25

func main() {

	input, err := readLinesToInt(inputPath)
	if err != nil {
		fmt.Println("Input read err: " + err.Error())
		return
	}

	var dataMap = make(map[int]int)
	var p2Target int

	for i := 0; i < len(input); i++ {
		dataMap[i] = input[i]
	}

	for i := preamble; i < len(input); i++ {
		var target = input[i]
		var good = false

		for j := i - preamble; j < i; j++ {
			for k := i - preamble; k < i; k++ {
				if j == k {
					break
				}
				if dataMap[j]+dataMap[k] == target {
					good = true
					break
				}
			}
			if good {
				break
			}
		}

		if !good {
			fmt.Printf("P1: %d\n", target)
			p2Target = target
			break
		}
	}

	for i := 0; i < len(input); i++ {
		tempScore, tempSmall, tempLarge := input[i], input[i], input[i]

		for j := i + 1; j < len(input); j++ {

			if input[j] < tempSmall {
				tempSmall = input[j]
			} else if input[j] > tempLarge {
				tempLarge = input[j]
			}

			tempScore += input[j]
			if tempScore == p2Target {
				fmt.Printf("P2: %d\n", tempSmall+tempLarge)
				break
			}
			if tempScore > p2Target {
				break
			}
		}
	}
}

// read lines of a file into a slice
func readLinesToInt(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Non int found in input!")
			break
		}

		lines = append(lines, num)
	}
	return lines, scanner.Err()
}
