package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var one = big.NewInt(1)

func crt(a, n []*big.Int) (*big.Int, error) {
	// No clue how this works, but it does.
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)

		// GCD, greates common denomenator is part of golang bigint
		z.GCD(nil, &s, n1, &q)

		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

func bussesPart2(s *bufio.Scanner) *big.Int {

	var rows []string

	for s.Scan() {
		rows = append(rows, strings.Trim(s.Text(), "\n"))
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	var a []*big.Int
	var n []*big.Int

	// Use chinese remainder theorem
	// https://rosettacode.org/wiki/Chinese_remainder_theorem#Go

	// Create a list of strings from the input
	for i, id := range strings.Split(rows[1], ",") {
		if id != "x" {
			ni, _ := strconv.Atoi(id) // ID's
			ai := ni - i
			a = append(a, big.NewInt(int64(ai)))
			n = append(n, big.NewInt(int64(ni)))
		}
	}

	ret, _ := crt(a, n) // run chines remainder theory
	return ret
}

func findEarliestBus(s *bufio.Scanner) int {
	var rows []string

	for s.Scan() {
		rows = append(rows, strings.Trim(s.Text(), "\n"))
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	// Parse arrival time
	target, _ := strconv.Atoi(rows[0])

	// answer should be large, we're looking for the smallest value
	answer := 10000000
	ret := 0

	for _, n := range strings.Split(rows[1], ",") {
		// do nothing with x
		if n == "x" {
			continue
		}

		// parse to int
		i, _ := strconv.Atoi(n)

		// set sum
		sum := i
		for {
			if sum > target {
				dt := sum - target
				if dt < answer {
					answer = dt
					ret = i * dt
				}
				break
			}
			sum += i
		}
	}

	return ret
}

func main() {
	part2 := true
	file, err := os.Open("data_13.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		n := findEarliestBus(scanner)
		fmt.Println("Day 12-1:", n)
	} else {
		// n := findSequentialBus(scanner, 100000000000000)
		n := bussesPart2(scanner)
		fmt.Println("Day 12-2:", n)
	}
}
