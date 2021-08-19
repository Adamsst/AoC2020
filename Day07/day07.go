package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputPath = "input.txt"

type Bag struct {
	Name     string
	GoesGold bool
	Members  []BagNameAndCount
}

type BagNameAndCount struct {
	Name  string
	Count int
}

type BagToCheck struct {
	Name       string
	Multiplier int
}

func main() {

	input, err := readLines(inputPath)
	if err != nil {
		fmt.Println("Input read err: " + err.Error())
		return
	}

	var bags = make(map[string]*Bag)

	for _, in := range input {
		var first = strings.Split(in, "contain")
		var key = strings.Split(first[0], "bags")

		bags[strings.TrimSpace(key[0])] = &Bag{
			Name:     strings.TrimSpace(key[0]),
			GoesGold: false,
		}

		first[1] = strings.TrimSpace(first[1])

		if first[1][0:2] == "no" {
			continue
		}

		addMembers(bags[strings.TrimSpace(key[0])], first[1])
	}

	var updated = true
	for updated {
		updated = false
		for key, _ := range bags {
			if key == "shiny gold" {
				continue
			}
			for _, m := range bags[key].Members {
				if m.Name == "shiny gold" {
					if !bags[key].GoesGold {
						bags[key].GoesGold = true
						updated = true
					}
					continue
				}

				if bags[m.Name].GoesGold {
					if !bags[key].GoesGold {
						bags[key].GoesGold = true
						updated = true
					}
				}
			}
		}
	}

	var score = 0
	for _, cur := range bags {
		if cur.GoesGold {
			score++
		}
	}

	fmt.Printf("P1: %d\n", score)

	var score2 = 0
	var bagsToCheck = make([]BagToCheck, 0)
	for _, cur := range bags["shiny gold"].Members {
		bagsToCheck = append(bagsToCheck, BagToCheck{
			Name:       cur.Name,
			Multiplier: cur.Count,
		})
		score2 += cur.Count
	}

	for i := 0; i < len(bagsToCheck); i++ {
		for _, cur := range bags[bagsToCheck[i].Name].Members {
			bagsToCheck = append(bagsToCheck, BagToCheck{
				Name:       cur.Name,
				Multiplier: (cur.Count * bagsToCheck[i].Multiplier),
			})
			score2 += (cur.Count * bagsToCheck[i].Multiplier)
		}
	}

	fmt.Printf("P2: %d\n", score2)
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

func addMembers(bag *Bag, containedBags string) {
	var curBagNum = 0
	bag.Members = make([]BagNameAndCount, 0)
	for len(containedBags) > 0 {
		bag.Members = append(bag.Members, BagNameAndCount{})
		numStr := strings.SplitN(containedBags, " ", 2)
		containedBags = numStr[1]
		num, err := strconv.Atoi(numStr[0])
		if err != nil {
			fmt.Println("uhh we got an err here trying to get the number of bags " + err.Error())
			return
		} else {
			bag.Members[curBagNum].Count = num
		}

		bagName := strings.SplitN(containedBags, "bag", 2)

		bag.Members[curBagNum].Name = strings.TrimSpace(bagName[0])

		if len(bagName) < 2 {
			return
		}

		if bagName[1][0] == 's' {
			bagName[1] = bagName[1][1:]
		}

		if bagName[1][0] == ',' {
			containedBags = strings.TrimSpace(bagName[1][2:])
		} else {
			containedBags = ""
		}
		curBagNum++
	}
}
