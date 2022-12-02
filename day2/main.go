package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Loose = 0
	Draw  = 3
	Win   = 6
)

var conditions = map[string]int8{
	"A": 1, // Rock
	"B": 2, // Paper
	"C": 3, // Scissors
	"X": 1, // Rock
	"Y": 2, // Paper
	"Z": 3, // Scissors
}

func main() {
	PartOne()
	PartTwo()
}

func RockPaperScissorResult(ourHand string, theirHand string) int {
	result := conditions[theirHand] - conditions[ourHand]
	if result == 0 {
		return Draw
	} else if result == -1 || result == 2 {
		return Win
	} else if result == 1 || result == -2 {
		return Loose
	}

	return 0
}

func PartOne() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	points := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hands := strings.Split(scanner.Text(), " ")
		theirHand := hands[0]
		ourHand := hands[1]

		points += int(conditions[ourHand])
		points += RockPaperScissorResult(ourHand, theirHand)
	}

	fmt.Println(points)
}

func PartTwo() {}
