package main

import (
	"bufio"
	"strings"
	"testing"
)

const example = `abc

a
b
c

ab
ac

a
a
a
a

b
`

func Test_countUniqueAnswers(t *testing.T) {

	r := strings.NewReader(example)
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
			name: "First Example",
			args: args{scanner},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countUniqueAnswers(tt.args.s); got != tt.want {
				t.Errorf("countUniqueAnswers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countCommonAnswers(t *testing.T) {

	r := strings.NewReader(example)
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
			name: "Second Example",
			args: args{scanner},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCommonAnswers(tt.args.s); got != tt.want {
				t.Errorf("countCommonAnswers() = %v, want %v", got, tt.want)
			}
		})
	}
}
