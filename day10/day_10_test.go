package main

import (
	"bufio"
	"strings"
	"testing"
)

const (
	example11 = `16
10
15
5
1
11
7
19
6
12
4
`
	example12 = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`
)

func createScanner(c string) *bufio.Scanner {
	r := strings.NewReader(c)
	return bufio.NewScanner(r)
}

func Test_joltageDifference(t *testing.T) {
	type args struct {
		s *bufio.Scanner
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1-1",
			args: args{createScanner(example11)},
			want: 35,
		},
		{
			name: "Example 1-2",
			args: args{createScanner(example12)},
			want: 220,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := joltageDifference(tt.args.s); got != tt.want {
				t.Errorf("joltageDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_adapterArangements(t *testing.T) {
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
			args: args{createScanner(example11)},
			want: 8,
		},
		// {
		// 	name: "Example 1-2",
		// 	args: args{createScanner(example12)},
		// 	want: 19208,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adapterArangements(tt.args.s); got != tt.want {
				t.Errorf("adapterArangements() = %v, want %v", got, tt.want)
			}
		})
	}
}
