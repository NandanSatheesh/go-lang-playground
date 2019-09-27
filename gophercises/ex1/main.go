package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Problem struct {
	Question string
	Answer   string
}

func main() {

	csvFileName := flag.String("csv", "D:\\go-lang-playground\\gophercises\\ex1\\problem.csv",
		"a csv in the format 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	checkErr(err)

	reader := csv.NewReader(file)
	lines, error := reader.ReadAll()
	checkErr(error)

	problems := parseLines(lines)
	counter := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d : %s \n", i+1, p.Question)

		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.Answer {
			counter++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", counter, len(problems))
}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))

	for i, line := range lines {
		ret[i] = Problem{
			Question: strings.TrimSpace(line[0]),
			Answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
