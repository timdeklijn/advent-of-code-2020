package main

import (
	"bufio"
	"strings"
	"testing"
)

const example1 = `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)
`

func createScanner(c string) *bufio.Scanner {
	r := strings.NewReader(c)
	return bufio.NewScanner(r)
}

func Test_checkAllergens(t *testing.T) {
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
			if got := checkAllergens(tt.args.s); got != tt.want {
				t.Errorf("checkAllergens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stockRaft(t *testing.T) {
	type args struct {
		s *bufio.Scanner
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{createScanner(example1)},
			want: "mxmxvkd,sqjhc,fvjkl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stockRaft(tt.args.s); got != tt.want {
				t.Errorf("stockRaft() = %v, want %v", got, tt.want)
			}
		})
	}
}
