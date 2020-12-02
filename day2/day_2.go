package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CheckPasswordsPart2 is the solution to day_2_2
func CheckPasswordsPart2(s bufio.Scanner) int {
	correctPW := 0
	for s.Scan() {
		// Split line on spaces
		splt := strings.Split(s.Text(), " ")

		// Get min and max occurances of character in password
		occSplit := strings.Split(splt[0], "-")
		pos1, err := strconv.Atoi(occSplit[0])
		if err != nil {
			panic(err)
		}
		pos2, err := strconv.Atoi(occSplit[1])
		if err != nil {
			panic(err)
		}

		// Get the character to count
		char := strings.Trim(splt[1], ":")

		// Check if char is in place
		ok1 := string(splt[2][pos1-1]) == char
		ok2 := string(splt[2][pos2-1]) == char

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
func CheckPasswords(s bufio.Scanner) int {
	correctPW := 0
	for s.Scan() {
		// Split line on spaces
		splt := strings.Split(s.Text(), " ")

		// Get min and max occurances of character in password
		occSplit := strings.Split(splt[0], "-")
		minOcc, err := strconv.Atoi(occSplit[0])
		if err != nil {
			panic(err)
		}
		maxOcc, err := strconv.Atoi(occSplit[1])
		if err != nil {
			panic(err)
		}

		// Get the character to count
		char := strings.Trim(splt[1], ":")

		// Count the character in the password
		charOcc := strings.Count(splt[2], char)

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
		fmt.Println("Day 1:", CheckPasswords(*scanner))
	}
	// Create file scanner
	file, err := os.Open("data_2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Run solution
	fmt.Println("Day 2:", CheckPasswordsPart2(*scanner))
}
