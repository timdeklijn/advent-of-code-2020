package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

const (
	example1 = `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12
`
	example2 = `class: 0-1 or 4-19
 row: 0-5 or 8-19
 seat: 0-13 or 16-19
 
 your ticket:
 11,12,13
 
 nearby tickets:
 3,9,18
 15,1,5
 5,14,9
 `
)

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
			want: 71,
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

func Test_decodeTicket(t *testing.T) {
	type args struct {
		s *bufio.Scanner
	}

	w := make(map[string]int)
	w["class"] = 12
	w["row"] = 11
	w["seat"] = 13

	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Example 2",
			args: args{createScanner(example2)},
			want: w,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeTicket(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeTicket() = %v, want %v", got, tt.want)
			}
		})
	}
}
