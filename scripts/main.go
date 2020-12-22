package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"time"
)

// Star containes the time in unix format it took a specific puzzle
type Star struct {
	GetStarTs string `json:"get_star_ts"`
}

// Member contains all data on a participant
type Member struct {
	CompletionDayLevel map[int]map[int]Star `json:"completion_day_level"`
	Stars              int                  `json:"start"`
	GlobalScore        int                  `json:"global_score"`
	LastStartTs        string               `json:"last_star_ts"`
	ID                 string               `json:"id"`
	LocalScore         int                  `json:"local_score"`
	Name               string               `json:"name"`
}

// Data is the top data structure in the JSON hierarchy
type Data struct {
	Members map[int]Member `json:"members"`
	OwnerID string         `json:"owner_id"`
	Event   string         `json:"event"`
}

// Output is a container for the data to output
type Output struct {
	day            int
	delta1, delta2 float64
}

// parsJSON parser the raw data file data.json to a Data structure
func parseJSON() Data {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	var data Data
	json.Unmarshal(byteValue, &data)

	return data
}

// parsData parses a Data structure to a []Output for member with number 886702
func parseData(data Data) []Output {
	var output []Output
	for dn, d := range data.Members[886702].CompletionDayLevel {

		// Parse unix times
		t1, err := strconv.ParseInt(d[1].GetStarTs, 10, 64)
		if err != nil {
			panic(err)
		}
		t2, err := strconv.ParseInt(d[2].GetStarTs, 10, 64)
		if err != nil {
			panic(err)
		}

		// Convert unix times to time deltas
		tm1 := time.Unix(t1, 0)
		tm2 := time.Unix(t2, 0)
		start := time.Date(tm1.Year(), tm1.Month(), tm1.Day(), 6, 0, 0, 0, tm1.Location())
		diff1 := tm1.Sub(start)
		diff2 := tm2.Sub(start)

		// Create Output
		output = append(output, Output{dn, diff1.Hours(), diff2.Hours()})
	}

	// Sort output on day number
	sort.Slice(output[:], func(i, j int) bool {
		return output[i].day < output[j].day
	})

	return output
}

// writeToFile writes the sorted []Output to a file called 'table'. The fields are
// separated by a " ".
func writeToFile(output []Output) {
	f, err := os.Create("table")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("day dt1 dt2\n")
	for _, d := range output {
		s := fmt.Sprintf("%d %f %f\n", d.day, d.delta1, d.delta2)
		f.WriteString(s)
	}
}

func main() {
	data := parseJSON()
	output := parseData(data)
	writeToFile(output)
}
