package main

import (
	"bufio"
	"math/big"
	"reflect"
	"strings"
	"testing"
)

const (
	example1 = `939
7,13,x,x,59,x,31,19
`
	example21 = `0
17,x,13,19
`
	example22 = `0
67,7,59,61
`
	example23 = `0
67,x,7,59,61
`
)

func createScanner(c string) *bufio.Scanner {
	r := strings.NewReader(c)
	return bufio.NewScanner(r)
}

func Test_findEarliestBus(t *testing.T) {
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
			want: 295,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findEarliestBus(tt.args.s); got != tt.want {
				t.Errorf("findEarliestBus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bussesPart2(t *testing.T) {
	type args struct {
		s *bufio.Scanner
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "Example 21",
			args: args{createScanner(example21)},
			want: big.NewInt(int64(3417)),
		},
		{
			name: "Example 22",
			args: args{createScanner(example22)},
			want: big.NewInt(int64(754018)),
		},
		{
			name: "Example 23",
			args: args{createScanner(example23)},
			want: big.NewInt(int64(779210)),
		},
		{
			name: "Example 2",
			args: args{createScanner(example1)},
			want: big.NewInt(int64(1068781)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bussesPart2(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bussesPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
