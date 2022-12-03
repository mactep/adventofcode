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

var theirHands = map[string]int8{
	"A": 1, // Rock
	"B": 2, // Paper
	"C": 3, // Scissors
}

var ourHands = map[string]int8{
	"X": 1, // Rock
	"Y": 2, // Paper
	"Z": 3, // Scissors
}

var outcomes = map[string]int8{
	"X": Loose,
	"Y": Draw,
	"Z": Win,
}

func main() {
	PartOne()
	PartTwo()
}

func RockPaperScissorResult(ourHand string, theirHand string) int {
	result := theirHands[theirHand] - theirHands[ourHand]
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

		points += int(ourHands[ourHand])
		points += RockPaperScissorResult(ourHand, theirHand)
	}

	fmt.Println(points)
}

func RockPaperScissorHand(expectedOutcome string, theirHand string) int {
	if expectedOutcome == "Y" {
		return int(theirHands[theirHand])
	} else if expectedOutcome == "X" {
		if theirHands[theirHand] == 1 {
			return int(theirHands[theirHand]) + 2
		} else {
			return int(theirHands[theirHand]) - 1
		}
	} else if expectedOutcome == "Z" {
		if theirHands[theirHand] == 3 {
			return int(theirHands[theirHand]) - 2
		} else {
			return int(theirHands[theirHand]) + 1
		}
	}

	return 0
}

func PartTwo() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	points := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hands := strings.Split(scanner.Text(), " ")
		theirHand := hands[0]
		expectedOutcome := hands[1]

		points += int(outcomes[expectedOutcome])
		points += RockPaperScissorHand(expectedOutcome, theirHand)
	}

	fmt.Println(points)
}
