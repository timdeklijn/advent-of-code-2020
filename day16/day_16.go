package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TicketField to save info per ticket
type TicketField struct {
	options []string
	pos     int
	nums    []int
	done    bool
}

func decodeTicket(s *bufio.Scanner) map[string]int {

	var allTickets [][]int
	var myTicketNums []int

	fields := make(map[string]map[int]bool)
	tickets := false
	myTicketFlag := false
	// sum := 0

	for s.Scan() {
		r := strings.Trim(s.Text(), "\n")

		if !tickets {

			// parse values
			if strings.Contains(r, " or ") {
				spl1 := strings.Split(r, ":")
				name := strings.Trim(spl1[0], " ")
				spl2 := strings.Split(strings.Trim(spl1[1], " "), " or ")

				fieldRange := make(map[int]bool)
				for _, spl := range spl2 {
					mm := strings.Split(spl, "-")
					min, _ := strconv.Atoi(mm[0])
					max, _ := strconv.Atoi(mm[1])
					for i := min; i < max+1; i++ {
						fieldRange[i] = true
					}
				}
				fields[name] = fieldRange
			}

			// Parse my Ticket
			if strings.Contains(r, "your ticket:") {
				myTicketFlag = true
				continue
			}

			if myTicketFlag {
				rr := strings.Split(r, ",")
				var tmp []int
				for _, n := range rr {
					nn, _ := strconv.Atoi(strings.Trim(n, " "))
					tmp = append(tmp, nn)
				}

				myTicketNums = tmp                   // save to later identify the fields
				allTickets = append(allTickets, tmp) // Can also be used to figure out rules
				myTicketFlag = false                 // Stop parsing my ticket
				continue
			}

			// We start parsing tickets
			if strings.Contains(r, "nearby tickets:") {
				tickets = true
			}

			continue
		}

		// parse other tickets

		rr := strings.Split(strings.Trim(r, " "), ",")
		if len(rr) > 1 { // Ugly, fix?
			var tmp []int
			for _, n := range rr {
				nn, _ := strconv.Atoi(n)
				tmp = append(tmp, nn)
			}
			allTickets = append(allTickets, tmp)
		}

	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	// Filter out the tickets that contain "wrong" numbers
	var correctTickets [][]int
	for i := 0; i < len(allTickets); i++ {
		cnt := 0
		for _, n := range allTickets[i] {

			for k := range fields {
				if ok := fields[k][n]; ok {
					cnt++
					break
				}
			}
		}
		if cnt == len(allTickets[0]) {
			correctTickets = append(correctTickets, allTickets[i])
		}
	}

	// Create a list of field names, we are going to eliminate fields from this
	var nameList []string
	for k := range fields {
		nameList = append(nameList, k)
	}

	// Create a list of ticketFields, these contain possible field names, column number
	// and all numbers in the column
	var ticketList []TicketField
	for i := 0; i < len(correctTickets[0]); i++ {
		var tmp []int
		for _, row := range correctTickets {
			tmp = append(tmp, row[i])
		}
		ticketList = append(ticketList, TicketField{nameList, i, tmp, false})
	}

	// for each column in ticketList
	for i := range ticketList {

		var newFields []string
		for _, field := range ticketList[i].options {

			// loop over all number and check if they all fit in a range
			save := true
			for _, n := range ticketList[i].nums {

				// if something does not fit, do not add the current field to newFields
				if _, ok := fields[field][n]; !ok {
					save = false
					break
				}

			}

			if save {
				newFields = append(newFields, field)
				continue
			}
		}
		ticketList[i].options = newFields
	}

	// Start eliminating
	for {
		exit := true         // exit
		var eliminate string // eliminate this one

		// Find an options list with a single element that is not "done" yet.
		for i := range ticketList {
			if len(ticketList[i].options) == 1 && !ticketList[i].done {
				eliminate = ticketList[i].options[0]
				ticketList[i].done = true
				exit = false // change
				break
			}
		}

		// Eliminate the "eliminate" option from optionlists that are still not of
		// length 1
		for i := range ticketList {
			var newOptions []string
			n := false
			for _, o := range ticketList[i].options {
				if len(ticketList[i].options) > 1 && o != eliminate {
					newOptions = append(newOptions, o)
					n = true
				}
			}
			if n {
				ticketList[i].options = newOptions
			}
		}

		// Break when needed
		if exit {
			break
		}
	}

	// Convert to usefull output
	myTicket := make(map[string]int)
	for _, v := range ticketList {
		myTicket[v.options[0]] = myTicketNums[v.pos]
	}

	return myTicket
}

func ticketScaningErrorRate(s *bufio.Scanner) int {

	rs := make(map[int]bool)
	tickets := false
	sum := 0

	for s.Scan() {
		r := strings.Trim(s.Text(), "\n")

		if !tickets {

			// parse values
			if strings.Contains(r, " or ") {
				spl1 := strings.Split(r, ":")
				spl2 := strings.Split(strings.Trim(spl1[1], " "), " or ")
				for _, spl := range spl2 {
					mm := strings.Split(spl, "-")
					min, _ := strconv.Atoi(mm[0])
					max, _ := strconv.Atoi(mm[1])
					for i := min; i < max+1; i++ {
						rs[i] = true
					}
				}
			}

			// We start parsing tickets
			if strings.Contains(r, "nearby tickets:") {
				tickets = true
			}

			continue
		}
		spl := strings.Split(r, ",")
		for _, n := range spl {
			nn, _ := strconv.Atoi(n)
			if _, ok := rs[nn]; !ok {
				sum += nn
			}
		}
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return sum
}

func main() {
	part2 := true
	file, err := os.Open("data_16.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !part2 {
		n := ticketScaningErrorRate(scanner)
		fmt.Println("Day 16-1:", n)
	} else {
		mp := decodeTicket(scanner)
		var l []int
		for k, v := range mp {
			if strings.Contains(k, "departure") {
				l = append(l, v)
			}
		}

		n := l[0]
		for i := 1; i < len(l); i++ {
			n *= l[i]
		}

		fmt.Println("Day 16-2:", n)
	}
}
