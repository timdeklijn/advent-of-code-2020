package main

import (
	"testing"
)

func Test_crabCupspt1(t *testing.T) {
	type args struct {
		input string
		turns int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1, 10 turns",
			args: args{"389125467", 10},
			want: 92658374,
		},
		{
			name: "Example 1, 100 turns",
			args: args{"389125467", 100},
			want: 67384529,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := crabCupspt1(tt.args.input, tt.args.turns); got != tt.want {
				t.Errorf("crabCupspt1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_crabCupspt2(t *testing.T) {
	type args struct {
		input string
		turns int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 2",
			args: args{"389125467", 10000000},
			want: 149245887792,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := crabCupspt2(tt.args.input, tt.args.turns); got != tt.want {
				t.Errorf("crabCupspt2() = %v, want %v", got, tt.want)
			}
		})
	}
}
