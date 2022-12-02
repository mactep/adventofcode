package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ParteTwo() {
	highest := [3]int{0, 0, 0}
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
			for i, high := range highest {
				if aggregated > high {
					highest[i] = aggregated
					break
				}
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

	fmt.Println(highest[0] + highest[1] + highest[2])
}
