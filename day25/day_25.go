package main

import (
	"fmt"
)

// Used to calculate the value
const DIVIDER = 20201227
const SUBJECT = 7

// getLoopSize returns the number of loops needed to get to a target value using the
// specified formula
func getLoopSize(target int) int {
	value := 1
	cnt := 0
	for value != target {
		value = value * SUBJECT % DIVIDER
		cnt++
	}
	return cnt
}

// getEncryption calculates the value using a key and a number of loops
func getEncryption(target, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value = value * target % DIVIDER
	}
	return value
}

// solvePart1 calculates the encryption key to the room
func solvePart1(n1, n2 int) int {
	loops1 := getLoopSize(n1)
	return getEncryption(n2, loops1)
}

func main() {
	n1 := 8252394
	n2 := 6269621
	fmt.Println("Day 25:", solvePart1(n1, n2))
}
