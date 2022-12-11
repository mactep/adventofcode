package main

import (
	"bufio"
	"strings"
	"testing"
)

const example = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestPartOne(t *testing.T) {
	tests := []struct {
		name  string
		input *bufio.Scanner
		want  int
	}{
		{
			name:  "Example 1",
			input: bufio.NewScanner(strings.NewReader(``)),
			want:  157,
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
		want  int
	}{
		{
			name:  "Example 2",
			input: bufio.NewScanner(strings.NewReader(``)),
			want:  70,
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
