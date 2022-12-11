package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func parseLine(line string) (int, int, int, int) {
	pattern := regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)
	matches := pattern.FindStringSubmatch(line)

	intValues := make([]int, 4)
	for i, match := range matches[1:] {
		intValues[i], _ = strconv.Atoi(match)
	}

	return intValues[0], intValues[1], intValues[2], intValues[3]
}

func isFullyContained(leftStart, leftEnd, rightStart, rightEnd int) bool {
	return leftStart >= rightStart && leftEnd <= rightEnd || rightStart >= leftStart && rightEnd <= leftEnd
}

func PartOne(scanner *bufio.Scanner) int {
	fullyContainedCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		leftStart, leftEnd, rightStart, rightEnd := parseLine(line)

		if isFullyContained(leftStart, leftEnd, rightStart, rightEnd) {
			fullyContainedCount++
		}
	}

	return fullyContainedCount
}

func isPartiallyContained(leftStart, leftEnd, rightStart, rightEnd int) bool {
	return !(leftEnd < rightStart || leftStart > rightEnd)
}

func PartTwo(scanner *bufio.Scanner) int {
	fullyContainedCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		leftStart, leftEnd, rightStart, rightEnd := parseLine(line)

		if isPartiallyContained(leftStart, leftEnd, rightStart, rightEnd) {
			fullyContainedCount++
		}
	}

	return fullyContainedCount
}
