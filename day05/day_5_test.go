package main

import (
	"reflect"
	"testing"
)

func Test_generateSequence(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test sequence",
			args: args{length: 5},
			want: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateSequence(tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarize(t *testing.T) {
	type args struct {
		s    string
		one  byte
		zero byte
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Binarize List",
			args: args{s: "aabbaa", one: 'a', zero: 'b'},
			want: []int{1, 1, 0, 0, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarize(tt.args.s, tt.args.one, tt.args.zero); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binarize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name:  "tst1",
			args:  args{s: "BFFFBBFRRR"},
			want:  70,
			want1: 7,
		},
		{
			name:  "tst2",
			args:  args{s: "FFFBBBFRRR"},
			want:  14,
			want1: 7,
		},
		{
			name:  "tst3",
			args:  args{s: "BBFFBBFRLL"},
			want:  102,
			want1: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Decode(tt.args.s)
			if got != tt.want {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Decode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
