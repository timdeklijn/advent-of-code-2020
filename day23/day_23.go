package main

import (
	"fmt"
	"strconv"
	"strings"
)

// length is used when looping over number
var length int

// Cup is a container for a linked list
type Cup struct {
	num  int
	next *Cup
}

// LinkedPickUp removes three cups from the linked list and returns them in a slice
func LinkedPickUp(c *Cup) []*Cup {

	// this will contain the taken out Cups
	var tmpList []*Cup

	// find the 4th cup after c. Append all cups to tmpList.
	next := c.next
	tmpList = append(tmpList, next)
	next = next.next
	tmpList = append(tmpList, next)
	next = next.next
	tmpList = append(tmpList, next)

	// link c to cup after final cup in tmplist
	c.next = next.next

	// unlink final cup in templist
	next.next = nil
	return tmpList
}

// printAll walks through the linked list and prints all nums in a string
func printAll(l map[int]*Cup) {
	s := ""
	currentCup := l[1]
	s += strconv.Itoa(currentCup.num)
	currentCup = currentCup.next

	for {
		if currentCup.num == 1 {
			break
		}
		s += strconv.Itoa(currentCup.num)
		currentCup = currentCup.next
	}
	fmt.Println(s)
}

// NewCupLinkedList creates a linked list of cups with nums taken from the input
// string
func NewBigCupLinkedList(s string) map[int]*Cup {
	// split string into numbers
	nums := strings.Split(s, "")

	// create first cup
	n1, _ := strconv.Atoi(nums[0])
	currentCup := &Cup{n1, nil}

	// loop over numbers and add cups
	l := make(map[int]*Cup)
	for _, s := range nums[1:] {
		n, _ := strconv.Atoi(s)   // convert input string to a number
		nextCup := &Cup{n, nil}   // create next cup
		currentCup.next = nextCup // link next cup from current cup
		l[currentCup.num] = currentCup
		currentCup = nextCup // repeat
	}

	// add cups with the number 10 - 1000000 to the linked list
	for i := 10; i <= 1000000; i++ {
		nextCup := &Cup{i, nil}
		currentCup.next = nextCup
		l[currentCup.num] = currentCup
		currentCup = nextCup
	}

	// add first cup as next to last cup (make a circle)
	currentCup.next = l[n1]
	l[currentCup.num] = currentCup

	// return list of cups
	return l
}

// NewCupLinkedList creates a linked list of cups with nums taken from the input
// string. The linked list is saved in a map.
func NewCupLinkedList(s string) map[int]*Cup {
	// split string into numbers
	nums := strings.Split(s, "")

	// create first cup
	n1, _ := strconv.Atoi(nums[0])
	currentCup := &Cup{n1, nil}

	// loop over numbers and link cups
	l := make(map[int]*Cup)
	for _, s := range nums[1:] {
		n, _ := strconv.Atoi(s)
		nextCup := &Cup{n, nil}
		// link the next cup to the current cup
		currentCup.next = nextCup
		l[currentCup.num] = currentCup
		currentCup = nextCup
	}
	// add first cup as next to last cup
	currentCup.next = l[n1]
	l[currentCup.num] = currentCup

	// return list of cups
	return l
}

// findLinkedDestinationCup finds the cup after which the taken out cups should be
// placed
func findLinkedDestinationCup(l map[int]*Cup, c *Cup, tmpList []*Cup) *Cup {
	target := c.num - 1
	for {
		if target == 0 {
			target = length
		}
		if target != tmpList[0].num && target != tmpList[1].num && target != tmpList[2].num && target != 0 {
			return l[target]
		}
		target--
	}
}

// linkedMove does one step in the game of cups using a linnked list
func linkedMove(l map[int]*Cup, c *Cup) *Cup {
	tmpList := LinkedPickUp(c)                             // take out list of three
	destination := findLinkedDestinationCup(l, c, tmpList) // find destination cup
	destinationNeighbour := destination.next               // save neighbour of destination cup
	destination.next = tmpList[0]                          // link destination cup to take out list
	tmpList[len(tmpList)-1].next = destinationNeighbour    // link take out list to desination neighbour
	return c.next                                          // return the next starting cup
}

// findLinkedAnswer ceates a string with all numbers from 1 until it sees 1 again.
func findLinkedAnswer(l map[int]*Cup) int {
	c := l[1]
	c = c.next
	s := ""

	for {
		if c.num == 1 {
			break
		}
		s += strconv.Itoa(c.num)
		c = c.next
	}
	n, _ := strconv.Atoi(s)
	return n
}

//findProduct takes the product of the two cup numbers after the cup with number 1
func findProduct(l map[int]*Cup) int {
	return l[1].next.num * l[1].next.next.num
}

// crabCupspt1 calculates the answer for part 1
func crabCupspt1(input string, turns int) int {
	cups := NewCupLinkedList(input)
	length = len(cups)
	startNum, _ := strconv.Atoi(string(input[0]))
	c := cups[startNum]
	for i := 0; i < turns; i++ {
		c = linkedMove(cups, c)
	}
	return findLinkedAnswer(cups)
}

// crabCupspt2 calculates the answer for part 2
func crabCupspt2(input string, turns int) int {
	// create cups
	cups := NewBigCupLinkedList(input)
	length = len(cups)
	startNum, _ := strconv.Atoi(string(input[0]))
	// first target cup
	c := cups[startNum]
	// simulate the turns
	for i := 0; i < turns; i++ {
		// do a single turn in the game
		c = linkedMove(cups, c)
	}
	// return the product of the two cups after cup with num 1
	return findProduct(cups)
}

func main() {
	part2 := true
	if !part2 {
		n := crabCupspt1("792845136", 100)
		fmt.Println("Day 22-1:", n)
	} else {
		n := crabCupspt2("792845136", 10000000)
		fmt.Println("Day 22-2:", n)
	}
}
