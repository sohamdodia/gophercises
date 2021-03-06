package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var filename = flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' ")
	flag.Parse()

	csvfile, err := os.Open(*filename)

	if err != nil {
		log.Fatalln("Could not open the csv file", err)
		os.Exit(1)
	}

	r := csv.NewReader(csvfile)

	lines, err := r.ReadAll()
	problems := parseLines(lines)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}
