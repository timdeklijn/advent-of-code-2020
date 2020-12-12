package main

import (
	"bufio"
	"strings"
	"testing"
)

const example1 = `F10
N3
F7
R90
F11
`

func createScanner(c string) *bufio.Scanner {
	r := strings.NewReader(c)
	return bufio.NewScanner(r)
}

func Test_calculatePosition(t *testing.T) {
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
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculatePosition(tt.args.s); got != tt.want {
				t.Errorf("calculatePosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculatePositionWithWaypoint(t *testing.T) {
	type args struct {
		s *bufio.Scanner
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 2",
			args: args{createScanner(example1)},
			want: 286,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculatePositionWithWaypoint(tt.args.s); got != tt.want {
				t.Errorf("calculatePositionWithWaypoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
