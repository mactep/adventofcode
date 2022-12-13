package main

import (
	"bufio"
	"strings"
	"testing"
)

const example = ``

func TestPartOne(t *testing.T) {
	tests := []struct {
		name  string
		input *bufio.Scanner
		want  int
	}{
		{
			name:  "Example 1",
			input: bufio.NewScanner(strings.NewReader(example)),
			want:  0,
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
			input: bufio.NewScanner(strings.NewReader(example)),
			want:  0,
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
