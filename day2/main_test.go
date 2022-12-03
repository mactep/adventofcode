package main

import (
	"io"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		name  string
		input io.Reader
		want  int
	}{
		{
			name:  "Example 1",
			input: strings.NewReader("A Y\nB X\nC Z"),
			want:  15,
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
		input io.Reader
		want  int
	}{
		{
			name:  "Example 2",
			input: strings.NewReader("A Y\nB X\nC Z"),
			want:  12,
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
