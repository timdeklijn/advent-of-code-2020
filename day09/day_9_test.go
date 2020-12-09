package main

import (
	"bufio"
	"strings"
	"testing"
)

const example1 = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`

func createScanner(c string) *bufio.Scanner {
	r := strings.NewReader(c)
	return bufio.NewScanner(r)
}

func Test_findWeakness(t *testing.T) {
	type args struct {
		s        *bufio.Scanner
		preamble int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Example 1",
			args:    args{createScanner(example1), 5},
			want:    127,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findWeakness(tt.args.s, tt.args.preamble)
			if (err != nil) != tt.wantErr {
				t.Errorf("findWeakness() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findWeakness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encryptionWeakness(t *testing.T) {
	type args struct {
		s        *bufio.Scanner
		preamble int
		target   int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Example 1",
			args:    args{createScanner(example1), 5, 127},
			want:    62,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encryptionWeakness(tt.args.s, tt.args.preamble, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("encryptionWeakness() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("encryptionWeakness() = %v, want %v", got, tt.want)
			}
		})
	}
}
