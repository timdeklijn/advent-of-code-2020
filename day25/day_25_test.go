package main

import (
	"testing"
)

func Test_getEncryption(t *testing.T) {
	type args struct {
		target   int
		loopSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Decrypt 1",
			args: args{5764801, 11},
			want: 14897079,
		},
		{
			name: "Decrypt 2",
			args: args{17807724, 8},
			want: 14897079,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEncryption(tt.args.target, tt.args.loopSize); got != tt.want {
				t.Errorf("getEncryption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLoopSize(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Loop Size 1",
			args: args{17807724},
			want: 11,
		},
		{
			name: "Loop Size 2",
			args: args{5764801},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLoopSize(tt.args.target); got != tt.want {
				t.Errorf("getLoopSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solvePart1(t *testing.T) {
	type args struct {
		n1 int
		n2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Solve Part 1",
			args: args{5764801, 17807724},
			want: 14897079,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart1(tt.args.n1, tt.args.n2); got != tt.want {
				t.Errorf("solvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}
