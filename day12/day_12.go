package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var heading = [4]string{"N", "E", "S", "W"}

// abs returns the absolute value of an integer
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// calcManhattenDist calculates the distance by simply summing the x and y coordinats
// (absolute values)
func calcManhattenDist(x, y int) int {
	return abs(x) + abs(y)
}

func calculatePosition(s *bufio.Scanner) int {
	pos := []int{0, 0} // keep track of the position
	headingIndex := 1  // We start heading east

	for s.Scan() {
		row := strings.Trim(s.Text(), "\n")
		com := string(row[0])
		n, _ := strconv.Atoi(row[1:])

		// Handle inputs
		if com == "L" {

			// New heading 'left', handle negative index
			headingIndex = headingIndex - (n / 90)
			if headingIndex < 0 {
				headingIndex += len(heading)
			}
			headingIndex = headingIndex % len(heading)
		}

		if com == "R" {
			headingIndex = (headingIndex + (n / 90)) % len(heading)
		}

		if com == "F" {
			// Change com to be the heading we are already facing
			com = heading[headingIndex]
		}

		if com == "N" {
			pos[0] += n
		}

		if com == "E" {
			pos[1] += n
		}

		if com == "S" {
			pos[0] -= n
		}

		if com == "W" {
			pos[1] -= n
		}

	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return calcManhattenDist(pos[0], pos[1])
}

func calculatePositionWithWaypoint(s *bufio.Scanner) int {
	shipPos := []int{0, 0} // keep track of the position of the ship
	wayPos := []int{1, 10} // the waypoint relative position starts at (1,10)

	for s.Scan() {
		// Parse a line of the input
		row := strings.Trim(s.Text(), "\n")
		com := string(row[0])
		n, _ := strconv.Atoi(row[1:])

		// Move ship towards waypoint by n steps
		if com == "F" {
			for i := 0; i < n; i++ {
				shipPos[0] += wayPos[0]
				shipPos[1] += wayPos[1]
			}
		}

		// Rotation to the right
		if com == "R" {
			// (1,10) -> (-10,1) -> (-1,-10) -> (10,-1)
			// Rotate n/90 times
			for i := 0; i < (n / 90); i++ {
				wayPos = []int{-wayPos[1], wayPos[0]}
			}
		}

		// Rotation to the left
		if com == "L" {
			// (1,10) -> (10,-1) -> (-1,-10) -> (-10,1)
			// Rotate n/90 times
			for i := 0; i < (n / 90); i++ {
				wayPos = []int{wayPos[1], -wayPos[0]}
			}
		}

		// Move waypoint based on com:

		if com == "N" {
			wayPos[0] += n
		}

		if com == "S" {
			wayPos[0] -= n
		}

		if com == "E" {
			wayPos[1] += n
		}

		if com == "W" {
			wayPos[1] -= n
		}

	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return calcManhattenDist(shipPos[0], shipPos[1])
}

func main() {
	part2 := true
	file, err := os.Open("data_12.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		n := calculatePosition(scanner)
		fmt.Println("Day 12-1:", n)
	} else {
		n := calculatePositionWithWaypoint(scanner)
		fmt.Println("Day 12-2:", n)
	}
}
