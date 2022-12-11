package main

import (
	"bufio"
	"fmt"
	"os"
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

func PartOne(scanner *bufio.Scanner) int {
	return 0
}

func PartTwo(scanner *bufio.Scanner) int {
	return 0
}
