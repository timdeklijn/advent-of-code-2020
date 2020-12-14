package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var adresses [][]string // Save processed adresses in

func createAdresses(o []string) {
	hasX := false  // Break out of recursion if we do not have X's
	var xIndex int // Save index of X

	// Look for X's
	for i, x := range o {
		if x == "X" {
			hasX = true
			xIndex = i
			break
		}
	}

	if hasX {
		// If we have an X copy o and modify it to be 0 or 1 and call createAdresses
		// again
		new0 := append([]string(nil), o...)
		new0[xIndex] = "0"
		createAdresses(new0)
		new1 := append([]string(nil), o...)
		new1[xIndex] = "1"
		createAdresses(new1)
	} else {
		// If we do not have X's, append current adress to adresses
		adresses = append(adresses, o)
	}
}

func maskAdress(s *bufio.Scanner) int {
	instructions := make(map[int]int) // Save adresses and values

	var mask []string

	for s.Scan() {
		r := strings.Trim(s.Text(), "\n")
		if r[:4] == "mask" {
			mask = strings.Split(r[7:], "")
		} else {
			// Parse adress + value
			spl := strings.Split(r, " ")
			value, _ := strconv.Atoi(spl[2])
			tmpAdress, _ := strconv.Atoi(
				strings.Trim(strings.Trim(spl[0], "mem["), "]"),
			)
			adress := int64(tmpAdress)

			// Convert adress to binary representation
			b := strings.Split(fmt.Sprintf("%036b", adress), "")

			// Modify values in adress to 1's or X's
			for i, m := range mask {
				switch m {
				case "1":
					b[i] = "1"
				case "X":
					b[i] = "X"
				default:
					continue
				}
			}

			// Recursively create all options for X's
			createAdresses(b)

			// Convert all options to ints and write value to adress
			for _, v := range adresses {
				n, _ := strconv.ParseInt(strings.Join(v[:], ""), 2, 64)
				instructions[int(n)] = value
			}
			// Clear adresses list
			adresses = adresses[:0]
		}
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	// Sum all values in instructions
	sum := 0
	for _, v := range instructions {
		sum += v
	}
	return sum
}

func executeInitProgram(s *bufio.Scanner) int64 {
	instructions := make(map[int]int64) // Save adresses and values

	var oneMask int64 = 0
	var zeroMask int64 = 0

	for s.Scan() {
		r := strings.Trim(s.Text(), "\n")

		// Parse mask
		if r[:4] == "mask" {

			// Create masks
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

			// mask the value
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
	part2 := true
	file, err := os.Open("data_14.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		n := executeInitProgram(scanner)
		fmt.Println("Day 14-1:", n)
	} else {
		n := maskAdress(scanner)
		fmt.Println("Day 14-2:", n)
	}
}
