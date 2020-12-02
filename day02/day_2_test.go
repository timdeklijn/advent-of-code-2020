package main

import (
	"bufio"
	"strings"
	"testing"
)

const s = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
`

func Test_checkPasswords(t *testing.T) {

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
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPasswords(tt.args.s); got != tt.want {
				t.Errorf("checkPasswords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPasswordsPart2(t *testing.T) {

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
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPasswordsPart2(tt.args.s); got != tt.want {
				t.Errorf("CheckPasswordsPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
