package main

import (
	"bufio"
	"strings"
	"testing"
)

const example1 = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
`

func createScanner(c string) *bufio.Scanner {
	r := strings.NewReader(c)
	return bufio.NewScanner(r)
}

func Test_playGame(t *testing.T) {
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
			want: 306,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := playGame(tt.args.s); got != tt.want {
				t.Errorf("playGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_playRecursiveGame(t *testing.T) {
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
			want: 291,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := playRecursiveGame(tt.args.s); got != tt.want {
				t.Errorf("playRecursiveGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
