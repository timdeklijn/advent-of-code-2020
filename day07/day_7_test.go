package main

import (
	"bufio"
	"strings"
	"testing"
)

const example1 = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
`

const example2 = `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.
`

const example3 = `shiny gold bags contain 2 dark red bags.
`

const example4 = `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
`

const example5 = `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
`

func Test_findParentBags(t *testing.T) {

	r := strings.NewReader(example1)
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
			name: "Example 1",
			args: args{scanner},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findParentBags(tt.args.s); got != tt.want {
				t.Errorf("findParentBags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createScanner(c string) *bufio.Scanner {
	r := strings.NewReader(c)
	return bufio.NewScanner(r)

}

func Test_countBagsInside(t *testing.T) {

	type args struct {
		s *bufio.Scanner
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 2",
			args: args{createScanner(example2)},
			want: 126,
		},
		{
			name: "Example 3",
			args: args{createScanner(example3)},
			want: 2,
		},
		{
			name: "Example 4",
			args: args{createScanner(example4)},
			want: 6,
		},
		{
			name: "Example 5",
			args: args{createScanner(example5)},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBagsInside(tt.args.s); got != tt.want {
				t.Errorf("countBagsInside() = %v, want %v", got, tt.want)
			}
		})
	}
}
