package main

import (
	"testing"
)

const tstData = `1721
979
366
299
675
1456`

func TestFindPair(t *testing.T) {
	type args struct {
		l []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "test case",
			args:    args{l: DataFromString(tstData)},
			want:    514579,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindPair(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindPair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindTriple(t *testing.T) {
	type args struct {
		l []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "test case",
			args:    args{l: DataFromString(tstData)},
			want:    241861950,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindTriple(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindTriple() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindTriple() = %v, want %v", got, tt.want)
			}
		})
	}
}
