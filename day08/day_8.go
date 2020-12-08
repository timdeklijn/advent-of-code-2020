package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Instruction saves the operations, the argument and the count
type Instruction struct {
	op    string
	arg   int
	count int
}

// parseScanner parses a scanner per line and returns a list of Instruction
func parseScanner(s *bufio.Scanner) []Instruction {
	var instructions []Instruction
	for s.Scan() {
		row := strings.Trim(s.Text(), "\n")
		spl := strings.Split(row, " ")
		n, _ := strconv.Atoi(spl[1])
		instructions = append(instructions, Instruction{spl[0], n, 0})
	}

	if err := s.Err(); err != nil {
		panic(err)
	}
	return instructions
}

// fixProg simply loops over the instructions changing nops to accs. Once we reach the
// end, we return the accumulator.
func fixProg(name string) int {
	changeInd := 0
	for {
		file, err := os.Open(name)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		instructions := parseScanner(scanner)
		if instructions[changeInd].op == "jmp" || instructions[changeInd].op == "nop" {
			if instructions[changeInd].op == "jmp" {
				instructions[changeInd].op = "nop"
			} else {
				instructions[changeInd].op = "acc"
			}
		}
		acc, ok := runInstructions(instructions, 100)
		if ok {
			return acc
		}
		changeInd++
	}
}

// runInstructions runs the input instructions. If a count is higher then lim we break
// and if we reach the last instriction we break.
func runInstructions(instructions []Instruction, lim int) (int, bool) {
	acc := 0
	ind := 0
	for {
		if ind == len(instructions)-1 {
			return acc, true
		}
		if instructions[ind].count > lim {
			return acc, false
		}
		switch instructions[ind].op {
		case "nop":
			instructions[ind].count++
			ind++
		case "acc":
			instructions[ind].count++
			acc += instructions[ind].arg
			ind++
		case "jmp":
			instructions[ind].count++
			ind += instructions[ind].arg
		}
	}
}

// calcAcc calculates the accumulator for the list of instructions in a Scanner.
func calcAcc(s *bufio.Scanner) int {
	acc, _ := runInstructions(parseScanner(s), 0)
	return acc
}

func main() {
	part2 := true
	file, err := os.Open("data_8.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		fmt.Println("Day 8-1:", calcAcc(scanner))
	} else {
		fmt.Println("Day 8-2:", fixProg("data_8.txt"))
	}
}
