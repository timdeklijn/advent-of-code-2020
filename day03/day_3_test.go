package main

import (
	"bufio"
	"strings"
	"testing"
)

const s = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func TestTreesInRoute(t *testing.T) {

	r := strings.NewReader(s)
	scanner := bufio.NewScanner(r)

	type args struct {
		s *bufio.Scanner
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example",
			args: args{
				s: scanner,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TreesInRoute(tt.args.s); got != tt.want {
				t.Errorf("TreesInRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreesInRouteProduct(t *testing.T) {

	r := strings.NewReader(s)
	scanner := bufio.NewScanner(r)

	type args struct {
		s *bufio.Scanner
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example",
			args: args{
				s: scanner,
			},
			want: 336,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TreesInRouteProduct(tt.args.s); got != tt.want {
				t.Errorf("TreesInRouteProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
