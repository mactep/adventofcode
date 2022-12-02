package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ParteOne() {
	highest := 0
	aggregated := 0

	const inputFile = "input"

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			if aggregated > highest {
				highest = aggregated
			}
			aggregated = 0
			continue
		}

		converted, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		aggregated += converted
	}

	fmt.Println(highest)
}
