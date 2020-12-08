package main

import (
	"bufio"
	"strings"
	"testing"
)

const example1 = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`

func createScanner(c string) *bufio.Scanner {
	r := strings.NewReader(c)
	return bufio.NewScanner(r)
}

func Test_calcAcc(t *testing.T) {
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
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcAcc(tt.args.s); got != tt.want {
				t.Errorf("calcAcc() = %v, want %v", got, tt.want)
			}
		})
	}
}
