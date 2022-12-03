package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println(PartOne(file))
	file.Seek(0, 0)
	fmt.Println(PartTwo(file))
}

func PartOne(file io.Reader) int {
	return 0
}

func PartTwo(file io.Reader) int {
	return 0
}
