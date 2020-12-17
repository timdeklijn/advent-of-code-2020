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

func Test_ticketScaningErrorRate(t *testing.T) {
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
			if got := ticketScaningErrorRate(tt.args.s); got != tt.want {
				t.Errorf("ticketScaningErrorRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
