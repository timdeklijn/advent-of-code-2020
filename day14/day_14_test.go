package main

import (
	"bufio"
	"strings"
	"testing"
)

const (
	example1 = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
`
	example2 = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
`

	example3 = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1
`
)

func createScanner(c string) *bufio.Scanner {
	r := strings.NewReader(c)
	return bufio.NewScanner(r)
}

func Test_executeInitProgram(t *testing.T) {
	type args struct {
		s *bufio.Scanner
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{createScanner(example1)},
			want: 165,
		},
		{
			name: "Example 2",
			args: args{createScanner(example2)},
			want: 165,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := executeInitProgram(tt.args.s); got != tt.want {
				t.Errorf("executeInitProgram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maskAdress(t *testing.T) {
	type args struct {
		s *bufio.Scanner
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 3",
			args: args{createScanner(example3)},
			want: 208,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maskAdress(tt.args.s); got != tt.want {
				t.Errorf("maskAdress() = %v, want %v", got, tt.want)
			}
		})
	}
}
