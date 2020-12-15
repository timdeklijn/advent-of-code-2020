package main

import "testing"

func Test_playGame(t *testing.T) {
	type args struct {
		in0 []int
		it  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{parseInput("0,3,6"), 2020},
			want: 436,
		},
		{
			name: "Example 2",
			args: args{parseInput("1,3,2"), 2020},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{parseInput("2,1,3"), 2020},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := playGame(tt.args.in0, tt.args.it); got != tt.want {
				t.Errorf("playGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
