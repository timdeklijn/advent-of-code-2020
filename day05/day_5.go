package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	// ROWS is the number of rows in the plane
	ROWS = 128
	// COLS is the number of columns in the plane
	COLS = 8
)

// generateSequence creates a int sequence starting at zero and of length length.
func generateSequence(length int) []int {
	var l []int
	for i := 0; i < length; i++ {
		l = append(l, i)
	}
	return l
}

// binarize converts a two character string into a list of 1's and 0's
func binarize(s string, one, zero byte) []int {
	var r []int
	for i := 0; i < len(s); i++ {
		switch c := s[i]; c {
		case one:
			r = append(r, 1)
		case zero:
			r = append(r, 0)
		}
	}
	return r
}

// Partition partitions a list based on a partition list by recursing over the partition
// list.
func Partition(part, rr []int) ([]int, []int) {
	if len(part) > 0 {
		switch d := part[0]; d {
		case 0:
			return Partition(part[1:], rr[:len(rr)/2])
		case 1:
			return Partition(part[1:], rr[len(rr)/2:])
		}
	}
	return part, rr
}

// Decode converts the binary search problem to a row and column value
func Decode(s string) (int, int) {
	rowString := s[:7]
	colString := s[6:]

	rowSeq := generateSequence(ROWS)
	rowPart := binarize(rowString, 'B', 'F')
	_, row := Partition(rowPart, rowSeq)

	colSeq := generateSequence(COLS)
	colPart := binarize(colString, 'R', 'L')
	_, col := Partition(colPart, colSeq)

	return row[0], col[0]
}

// HighestID will check all ID's for the highest one
func HighestID(s *bufio.Scanner) int {
	highest := 0
	for s.Scan() {
		row := strings.Trim(s.Text(), "\n")
		r, c := Decode(row)
		id := r*8 + c
		if id > highest {
			highest = id
		}
	}

	if err := s.Err(); err != nil {
		panic(err)
	}
	return highest
}

// FindSeat will find the only missing seat using ID's
func FindSeat(s *bufio.Scanner) int {
	// Add all id's to a list
	var seats []int
	for s.Scan() {
		row := strings.Trim(s.Text(), "\n")
		r, c := Decode(row)
		seats = append(seats, r*8+c)
	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	// Sort the list so we can find gaps larger then 1
	sort.Ints(seats)
	for i := 1; i < len(seats)-1; i++ {
		if seats[i]-seats[i-1] != 1 {
			return seats[i] - 1
		}
	}
	return 0
}

func main() {
	part2 := true
	file, err := os.Open("data_5.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		fmt.Println("Day 5-1:", HighestID(scanner))
	} else {
		fmt.Println("Day 5-2:", FindSeat(scanner))
	}
}
