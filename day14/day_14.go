package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func executeInitProgram(s *bufio.Scanner) int64 {
	instructions := make(map[int]int64) // Save adresses and values

	var oneMask int64 = 0
	var zeroMask int64 = 0

	for s.Scan() {
		r := strings.Trim(s.Text(), "\n")

		// Parse mask
		if r[:4] == "mask" {

			oneMask, _ = strconv.ParseInt(strings.ReplaceAll(r[7:], "X", "1"), 2, 0)
			zeroMask, _ = strconv.ParseInt(strings.ReplaceAll(r[7:], "X", "0"), 2, 0)

		} else {
			// Parse adress + value
			spl := strings.Split(r, " ")
			tmpValue, _ := strconv.Atoi(spl[2])
			value := int64(tmpValue)
			adress, _ := strconv.Atoi(
				strings.Trim(strings.Trim(spl[0], "mem["), "]"),
			)

			value &= oneMask
			value |= zeroMask

			// Save to adress
			instructions[adress] = value
		}
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	// Sum all instruction values
	var sum int64 = 0
	for _, v := range instructions {
		sum += v
	}

	return sum
}

func main() {
	part2 := false
	file, err := os.Open("data_14.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		n := executeInitProgram(scanner)
		fmt.Println("Day 14-1:", n)
	}
	// else {
	// 	// n := findSequentialBus(scanner)
	// 	n := bussesPart2(scanner)
	// 	fmt.Println("Day 14-2:", n)
	// }
}
