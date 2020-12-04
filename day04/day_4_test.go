package main

import (
	"bufio"
	"strings"
	"testing"
)

const s = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in

`

const valid = `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719

`

const invalid = `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007

`

func TestCheckPassports(t *testing.T) {

	r := strings.NewReader(s)
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
			name: "Example",
			args: args{scanner},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPassportsPart1(tt.args.s); got != tt.want {
				t.Errorf("CheckPassports() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCorrect(t *testing.T) {
	type args struct {
		l map[string]string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Complete",
			args: args{
				map[string]string{
					"byr": "ok",
					"iyr": "ok",
					"eyr": "ok",
					"hgt": "ok",
					"hcl": "ok",
					"ecl": "ok",
					"pid": "ok",
				},
			},
			want: true,
		},
		{
			name: "Incomplete",
			args: args{
				map[string]string{
					"eyr": "ok",
					"hgt": "ok",
					"hcl": "ok",
					"ecl": "ok",
					"pid": "ok",
				},
			},
			want: false,
		},
		{
			name: "Empty",
			args: args{make(map[string]string)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCorrect(tt.args.l); got != tt.want {
				t.Errorf("IsCorrect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPassportsPart2(t *testing.T) {

	r := strings.NewReader(valid)
	scanner := bufio.NewScanner(r)

	r2 := strings.NewReader(invalid)
	scanner2 := bufio.NewScanner(r2)

	type args struct {
		s *bufio.Scanner
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example",
			args: args{scanner},
			want: 4,
		},
		{
			name: "Example",
			args: args{scanner2},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPassportsPart2(tt.args.s); got != tt.want {
				t.Errorf("CheckPassportsPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkByr(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "OK",
			args: args{"2002"},
			want: true,
		},
		{
			name: "OK",
			args: args{"2003"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkByr(tt.args.v); got != tt.want {
				t.Errorf("checkByr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkHgt(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "v1",
			args: args{"60in"},
			want: true,
		},
		{
			name: "v2",
			args: args{"190cm"},
			want: true,
		},
		{
			name: "i1",
			args: args{"190in"},
			want: false,
		},
		{
			name: "i2",
			args: args{"190"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkHgt(tt.args.v); got != tt.want {
				t.Errorf("checkHgt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkHcl(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "v1",
			args: args{"#123abc"},
			want: true,
		},
		{
			name: "i1",
			args: args{"#123abz"},
			want: false,
		},
		{
			name: "i1",
			args: args{"123abc"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkHcl(tt.args.v); got != tt.want {
				t.Errorf("checkHcl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkEcl(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "v1",
			args: args{"brn"},
			want: true,
		},
		{
			name: "i1",
			args: args{"wat"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkEcl(tt.args.v); got != tt.want {
				t.Errorf("checkEcl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPid(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "v1",
			args: args{"000000001"},
			want: true,
		},
		{
			name: "i1",
			args: args{"0123456789"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPid(tt.args.v); got != tt.want {
				t.Errorf("checkPid() = %v, want %v", got, tt.want)
			}
		})
	}
}
