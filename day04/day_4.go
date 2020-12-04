package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// =====================================================================================
// CHECK FIELDS
// =====================================================================================

var fields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

// IsCorrect checks if all fields are present in a map
func IsCorrect(l map[string]string) bool {
	for _, f := range fields {
		_, found := l[f]
		if !found {
			return false
		}
	}
	return true
}

// =====================================================================================
// DATA VERIFICATION FUNCTIONS
// =====================================================================================

func checkByr(v string) bool {
	i, _ := strconv.Atoi(v)
	if i < 1920 || i > 2002 {
		return false
	}
	return true
}

func checkIyr(v string) bool {
	i, _ := strconv.Atoi(v)
	if i < 2010 || i > 2020 {
		return false
	}
	return true
}

func checkEyr(v string) bool {
	i, _ := strconv.Atoi(v)
	if i < 2020 || i > 2030 {
		return false
	}
	return true
}

func checkHgt(v string) bool {
	unit := v[len(v)-2:]
	if unit == "cm" || unit == "in" {
		i, _ := strconv.Atoi(v[:len(v)-2])
		if unit == "cm" {
			if i < 150 || i > 193 {
				return false
			}
		}
		if unit == "in" {
			if i < 59 || i > 76 {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func checkHcl(v string) bool {
	if v[0] != '#' {
		return false
	}
	for i := 1; i < len(v); i++ {
		el := v[i : i+1]
		if _, err := strconv.Atoi(el); err != nil {
			ok := false
			for _, opt := range []string{"a", "b", "c", "d", "e", "f"} {
				if el == opt {
					ok = true
				}
			}
			if !ok {
				return false
			}
		}
	}
	return true
}

func checkPid(v string) bool {
	if len(v) != 9 {
		return false
	}
	if _, err := strconv.Atoi(v); err != nil {
		return false
	}
	return true
}

func checkEcl(v string) bool {
	ok := false
	for _, opt := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if v == opt {
			ok = true
		}
	}
	if !ok {
		return false
	}
	return true
}

// HasCorrectData verifies the data in a passport
func HasCorrectData(l map[string]string) bool {
	for k, v := range l {
		switch k {
		case "byr":
			if !checkByr(v) {
				return false
			}
		case "iyr":
			if !checkIyr(v) {
				return false
			}
		case "eyr":
			if !checkEyr(v) {
				return false
			}
		case "hgt":
			if !checkHgt(v) {
				return false
			}
		case "hcl":
			if !checkHcl(v) {
				return false
			}
		case "ecl":
			if !checkEcl(v) {
				return false
			}
		case "pid":
			if !checkPid(v) {
				return false
			}
		}
	}
	return true
}

// =====================================================================================
// SOLUTIONS
// =====================================================================================

// CheckPassportsPart2 is the solution to part 2
func CheckPassportsPart2(s *bufio.Scanner) int {
	validCount := 0
	currentPW := make(map[string]string)
	for s.Scan() {
		row := strings.Trim(s.Text(), "\n")

		// Check password
		if len(row) == 0 {
			if IsCorrect(currentPW) && HasCorrectData(currentPW) {
				validCount++
			}

			currentPW = make(map[string]string)
		} else {
			// Fill in key value pairs into new passport
			splits := strings.Split(row, " ")
			for _, ss := range splits {
				k := strings.Split(ss, ":")
				currentPW[k[0]] = k[1]
			}
		}
	}

	// Check final passport
	if IsCorrect(currentPW) && HasCorrectData(currentPW) {
		validCount++
	}

	if err := s.Err(); err != nil {
		panic(err)
	}
	return validCount
}

// CheckPassportsPart1 is the solution to part 1
func CheckPassportsPart1(s *bufio.Scanner) int {
	validCount := 0
	currentPW := make(map[string]string)
	for s.Scan() {
		row := strings.Trim(s.Text(), "\n")

		if len(row) == 0 {
			if IsCorrect(currentPW) {
				validCount++
			}

			currentPW = make(map[string]string)
		} else {
			splits := strings.Split(row, " ")
			for _, ss := range splits {
				k := strings.Split(ss, ":")
				currentPW[k[0]] = k[1]
			}
		}
	}
	if IsCorrect(currentPW) {
		validCount++
	}

	if err := s.Err(); err != nil {
		panic(err)
	}
	return validCount
}

func main() {
	part2 := true
	file, err := os.Open("data_4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		fmt.Println("Day 4-1:", CheckPassportsPart1(scanner))
	} else {
		fmt.Println("Day 4-2:", CheckPassportsPart2(scanner))
	}
}
