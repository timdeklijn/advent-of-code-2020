package main

import (
	"bufio"
	"strings"
	"testing"
)

const (
	example1 = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`
	example2 = `.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....
`

	example3 = `.............
.L.L.#.#.#.#.
.............
`
)

func createScanner(c string) *bufio.Scanner {
	r := strings.NewReader(c)
	return bufio.NewScanner(r)
}

func Test_modelSeats(t *testing.T) {
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
			want: 37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modelSeats(tt.args.s); got != tt.want {
				t.Errorf("modelSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_modelSeatsDistance(t *testing.T) {
	type args struct {
		s *bufio.Scanner
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 2-1",
			args: args{createScanner(example2)},
			want: 8,
		},
		{
			name: "Example 2",
			args: args{createScanner(example1)},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modelSeatsDistance(tt.args.s); got != tt.want {
				t.Errorf("modelSeatsDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
