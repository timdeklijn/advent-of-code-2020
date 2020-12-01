package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ScannerToList converts a filescanner to a list of integers
func ScannerToList(scanner *bufio.Scanner) []int {
	var l []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		l = append(l, i)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return l
}

// DataFromFile creates a scanner from file and convert that to a list
// of integers.
func DataFromFile(name string) []int {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	return ScannerToList(scanner)
}

// DataFromString creates a reader from a string, a scanner from the reader and a
// list of integers from the scanner.
func DataFromString(s string) []int {
	r := strings.NewReader(s)
	scanner := bufio.NewScanner(r)
	return ScannerToList(scanner)

}

// FindPair return the product of two elements of a list of integers that
// sum up to 2020
func FindPair(l []int) (int, error) {
	for _, i := range l {
		for _, j := range l {
			if i+j == 2020 {
				return i * j, nil
			}
		}
	}
	return 0, fmt.Errorf("error")
}

// FindTriple return the product of three elements of a list of integers that
// sum up to 2020
func FindTriple(l []int) (int, error) {
	for _, i := range l {
		for _, j := range l {
			for _, k := range l {
				if i+j+k == 2020 {
					return i * j * k, nil
				}
			}
		}
	}
	return 0, fmt.Errorf("error")
}

func main() {
	fmt.Println("Day 1")
	// Do part 1 or part 2
	part2 := true
	if !part2 {
		// Choose dataloader
		l := DataFromFile
		// Find the product of the two elements in the list that sum to 2020
		r, err := FindPair(l("data_1.txt"))
		if err != nil {
			panic(err)
		}
		fmt.Println(r)
	} else {
		l := DataFromFile
		r, err := FindTriple(l("data_1.txt"))
		if err != nil {
			panic(err)
		}
		fmt.Println(r)
	}
}
