package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Problem struct {
	Question string
	Answer   string
}

func main() {

	csvFileName := flag.String("csv", "D:\\go-lang-playground\\gophercises\\ex1\\problem.csv",
		"a csv in the format 'question,answer'")
	timeLimit := flag.Int("limit", 3, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	checkErr(err)

	reader := csv.NewReader(file)
	lines, error := reader.ReadAll()
	checkErr(error)

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	counter := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d : %s \n", i+1, p.Question)
		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", counter, len(problems))
			return
		case answer := <-answerChannel:
			if answer == p.Answer {
				counter++
			}
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
