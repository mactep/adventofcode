package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

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

func ParseStacksLine(line string, stacks []structures.Stack) (end bool) {
	end = true
	line = line + " "
	cratePattern := regexp.MustCompile(`\[([A-Z])\]`)
	for i := 0; i+3 < len(line); i += 4 {
		matches := cratePattern.FindStringSubmatch(line[i : i+3])
		if len(matches) > 0 {
			stacks[i/4].PushBottom(matches[1])
			end = false
		}
	}

	return end
}

func ParseStacks(scanner *bufio.Scanner) (stacks []structures.Stack) {
	scanner.Scan()
	line := scanner.Text()
	stacks = make([]structures.Stack, (len(line)+1)/4)
	for i := 0; i < len(stacks); i++ {
		stacks[i] = structures.NewStack()

	}

	ParseStacksLine(line, stacks)

	for scanner.Scan() {
		line := scanner.Text()
		if ParseStacksLine(line, stacks) {
			break
		}
	}

	return stacks
}

func GetTopCrates(stacks []structures.Stack) string {
	topCrates := ""
	for _, stack := range stacks {
		topCrate := stack.Peek()
		if topCrate != nil {
			str, ok := topCrate.(string)
			if ok {
				topCrates += str
			}
		}
	}
	return topCrates
}

func ParseMoveLine(line string) (int, int, int, error) {
	movePattern := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	matches := movePattern.FindStringSubmatch(line)
	if len(matches) > 0 {
		amount, _ := strconv.Atoi(matches[1])
		from, _ := strconv.Atoi(matches[2])
		to, _ := strconv.Atoi(matches[3])
		return amount, from, to, nil
	}

	return 0, 0, 0, fmt.Errorf("Invalid move line: %s", line)
}

func PartOne(scanner *bufio.Scanner) string {
	stacks := ParseStacks(scanner)

	// ignore the blank line
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		amount, from, to, err := ParseMoveLine(line)
		if err != nil {
			panic(err)
		}

		for i := 0; i < amount; i++ {
			movedCrate := stacks[from-1].Pop()
			stacks[to-1].Push(movedCrate)
		}
	}

	topCrates := GetTopCrates(stacks)

	return topCrates
}

func PartTwo(scanner *bufio.Scanner) string {
	stacks := ParseStacks(scanner)

	// ignore the blank line
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		amount, from, to, err := ParseMoveLine(line)
		if err != nil {
			panic(err)
		}

		movingCrates := make([]string, amount)
		for i := 0; i < amount; i++ {
			movingCrates[i] = stacks[from-1].Pop().(string)
		}

		for i := amount - 1; i >= 0; i-- {
			stacks[to-1].Push(movingCrates[i])
		}
	}

	topCrates := GetTopCrates(stacks)

	return topCrates
}
