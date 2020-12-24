package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

const example1 = `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew
`

func createScanner(c string) *bufio.Scanner {
	r := strings.NewReader(c)
	return bufio.NewScanner(r)
}

func Test_tileFloor(t *testing.T) {
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
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tileFloor(tt.args.s); got != tt.want {
				t.Errorf("tileFloor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matchToCoordinate(t *testing.T) {
	type args struct {
		l []string
	}
	tests := []struct {
		name string
		args args
		want Coord
	}{
		{
			name: "test coordinates",
			args: args{[]string{"nw", "w", "sw", "e", "e"}},
			want: Coord{0, 0},
		},
		{
			name: "test coordinates",
			args: args{[]string{"e", "se", "w"}},
			want: Coord{1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchToCoordinate(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matchToCoordinate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_simulateTiles(t *testing.T) {
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
			args: args{createScanner(example1)},
			want: 2208,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simulateTiles(tt.args.s); got != tt.want {
				t.Errorf("simulateTiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
