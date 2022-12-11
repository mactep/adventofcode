package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mactep/aoc2022/structures"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Println(PartOne(scanner))
	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	fmt.Println(PartTwo(scanner))
}

func ItemPriority(item string) int {
	if item[0] > 'a' {
		return int(item[0]) - 'a' + 1
	}
	return int(item[0]) - 'A' + 26 + 1
}

func PartOne(scanner *bufio.Scanner) int {
	prioritySum := 0

	for scanner.Scan() {
		firstHalfItems := structures.NewSet()
		line := scanner.Text()
		for i := 0; i < len(line)/2; i++ {
			firstHalfItems.Add(string(line[i]))
		}

		secondHalfItems := structures.NewSet()
		for i := (len(line) / 2); i < len(line); i++ {
			secondHalfItems.Add(string(line[i]))
		}

		intersection := firstHalfItems.Intersection(secondHalfItems)
		for value := range intersection {
			prioritySum += ItemPriority(value)
		}
	}

	return prioritySum
}

func PartTwo(scanner *bufio.Scanner) int {
	prioritySum := 0

	for true {
		sets := make([]structures.Set, 3)
		for i := 0; i < 3; i++ {
			sets[i] = structures.NewSet()
			if !scanner.Scan() {
				return prioritySum
			}

			line := scanner.Text()
			for item := range line {
				sets[i].Add(string(line[item]))
			}
		}

		intersection := sets[0].Intersection(sets[1])
		intersection = intersection.Intersection(sets[2])
		for value := range intersection {
			prioritySum += ItemPriority(value)
		}
	}

	return prioritySum
}
