package main

import (
	"bufio"
	"strings"
	"testing"
)

func createScanner(c string) *bufio.Scanner {
	r := strings.NewReader(c)
	return bufio.NewScanner(r)
}

func Test_homeworkPart1(t *testing.T) {
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
			args: args{createScanner("1 + 2 * 3 + 4 * 5 + 6")},
			want: 71,
		},
		{
			name: "Example 1-2",
			args: args{createScanner("1 + (2 * 3) + (4 * (5 + 6))")},
			want: 51,
		},
		{
			name: "Example 1-3",
			args: args{createScanner("5 + (8 * 3 + 9 + 3 * 4 * 3)")},
			want: 437,
		},
		{
			name: "Example 1-4",
			args: args{createScanner("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))")},
			want: 12240,
		},
		{
			name: "Example 1-4",
			args: args{createScanner("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2")},
			want: 13632,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := homeworkPart1(tt.args.s); got != tt.want {
				t.Errorf("homeworkPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}
