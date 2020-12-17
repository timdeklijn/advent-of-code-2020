package main

import (
	"bufio"
	"strings"
	"testing"
)

const example1 = `.#.
..#
###
`

func createScanner(c string) *bufio.Scanner {
	r := strings.NewReader(c)
	return bufio.NewScanner(r)
}

func Test_simulateGrid(t *testing.T) {
	type args struct {
		s *bufio.Scanner
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{createScanner(example1)},
			want: 112,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simulateGrid(tt.args.s); got != tt.want {
				t.Errorf("simulateGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_simulateGrid4(t *testing.T) {
	type args struct {
		s *bufio.Scanner
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{createScanner(example1)},
			want: 848,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simulateGrid4(tt.args.s); got != tt.want {
				t.Errorf("simulateGrid4() = %v, want %v", got, tt.want)
			}
		})
	}
}
