package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

const filePath = "/Users/nandan/Documents/go-lang-playground/gophercises/ex1/problem.csv"

type Problem struct {
	question string
	answer   string
}

func main() {
	csvFileName := flag.String("csv", filePath, "A csv file which contains all problems")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	checkErr(err)

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	checkErr(err)

	Problems := make([]Problem, len(lines))

	for i, line := range lines {
		Problems[i] = Problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	count := 0

	for i, p := range Problems {
		fmt.Printf("Problem #%d: %s = \n", i, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == p.answer {
			count++
		}
	}

	fmt.Printf("Your score is %d", count)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
