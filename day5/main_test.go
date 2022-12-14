package main

import (
	"bufio"
	"strings"
	"testing"
)

const example = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestPartOne(t *testing.T) {
	tests := []struct {
		name  string
		input *bufio.Scanner
		want  string
	}{
		{
			name:  "Example 1",
			input: bufio.NewScanner(strings.NewReader(example)),
			want:  "CMZ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartOne(tt.input); got != tt.want {
				t.Errorf("PartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		name  string
		input *bufio.Scanner
		want  string
	}{
		{
			name:  "Example 2",
			input: bufio.NewScanner(strings.NewReader(example)),
			want:  "MCD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartTwo(tt.input); got != tt.want {
				t.Errorf("PartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
