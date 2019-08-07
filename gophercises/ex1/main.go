package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var Questions []string
var Answers []int64

func init() {

	file, err := os.Open("problem.csv")
	checkErr(err)

	csvData := csv.NewReader(file)

	for {
		row, err := csvData.Read()

		if err == io.EOF {
			break
		}
		Questions = append(Questions, row[0])
		number, _ := strconv.ParseInt(row[1], 10, 64)
		Answers = append(Answers, number)
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	score := 0
	for i := 0; i < len(Questions); i++ {

		fmt.Println("What's ", Questions[i], "?")

		inputText, err := reader.ReadString('\n')
		inputText = strings.Replace(inputText, "\n", "", -1)
		checkErr(err)

		userAnswer, err := strconv.ParseInt(inputText, 10, 64)
		checkErr(err)

		if userAnswer == Answers[i] {
			score++
		}

	}

	fmt.Println("Your score is ", score, ".")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
