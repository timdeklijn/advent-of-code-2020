package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Parse line to int-int-string-string removing some special chars
func lineSplitter(s string) (int, int, string, string) {
	// Split line on spaces
	splt := strings.Split(s, " ")

	// Parse ints in the fist element of the split list
	spl1 := strings.Split(splt[0], "-")
	i1, err := strconv.Atoi(spl1[0])
	if err != nil {
		panic(err)
	}
	i2, err := strconv.Atoi(spl1[1])
	if err != nil {
		panic(err)
	}

	// Remove ':' from second element of the split
	char := strings.Trim(splt[1], ":")
	return i1, i2, char, splt[2]
}

// CheckPasswordsPart2 is the solution to day_2_2
func CheckPasswordsPart2(s *bufio.Scanner) int {
	correctPW := 0
	for s.Scan() {

		// parse line
		pos1, pos2, char, pw := lineSplitter(s.Text())

		// Check if char is in place
		ok1 := string(pw[pos1-1]) == char
		ok2 := string(pw[pos2-1]) == char

		// There should be exactly 1 true. So that is an XOR operation.
		if (ok1 || ok2) && !(ok1 && ok2) {
			correctPW++
		}

	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	return correctPW
}

// CheckPasswords is the solution to day_2_1
func CheckPasswords(s *bufio.Scanner) int {
	correctPW := 0
	for s.Scan() {

		// parse line
		minOcc, maxOcc, char, pw := lineSplitter(s.Text())

		// Count the character in the password
		charOcc := strings.Count(pw, char)

		// Assert if password is correct
		if charOcc >= minOcc && charOcc <= maxOcc {
			correctPW++
		}

	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	return correctPW
}

func main() {
	day1 := false
	if day1 {
		// Create file scanner
		file, err := os.Open("data_2.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)

		// Run solution
		fmt.Println("Day 1:", CheckPasswords(scanner))
	}
	// Create file scanner
	file, err := os.Open("data_2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Run solution
	fmt.Println("Day 2:", CheckPasswordsPart2(scanner))
}
