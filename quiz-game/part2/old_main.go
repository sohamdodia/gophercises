package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"time"
)

func main() {
	var filename = flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' ")
	var timelimit = flag.Int("limit", 30, "a time limit for the quiz in seconds")

	flag.Parse()

	csvfile, err := os.Open(*filename)

	if err != nil {
		log.Fatalln("Could not open the csv file", err)
		os.Exit(1)
	}

	r := csv.NewReader(csvfile)

	lines, err := r.ReadAll()
	problems := parseLines(lines)
	rand.Seed(time.Now().UnixNano()) // do it once during app initialization
	Shuffle(problems)
	correct := 0

	fmt.Printf("Please hit enter to start the timer")
	var abc string
	fmt.Scanf("%s", &abc)
	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)

	defer timer.Stop()
	go func() {
		<-timer.C
		fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
		os.Exit(1)
	}()

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

func Shuffle(slice interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()
	for i := length - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		swap(i, j)
	}
}
