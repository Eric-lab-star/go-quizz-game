package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Problems struct {
	question string
	answer   string
}

func main() {

	csvFileName := flag.String("csv", "problems.csv", "a name of the csv file formated 'question,anser'")
	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("failed to open file name: %v. ", *csvFileName))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		exit("faild to parse file")
	}
	problems := parsedProblems(lines)
	score := 0
	timer := time.NewTimer(time.Duration(2) * time.Second)

	for index, value := range problems {
		fmt.Printf("Problem #%d: %s= ", index+1, value.question)
		answerCh := make(chan string)
		var answer string
		go Quiz(answerCh)
		select {
		case <-timer.C:
			fmt.Printf("\nTime Over \nToal Score: %v\n", score)
			return
		case answer = <-answerCh:
			if answer != value.answer {
				fmt.Println("inCorrect")
			} else {
				score++
				fmt.Printf("Correct!\ncurrent score: %v \n", score)
			}
		}

	}
	fmt.Printf("Total score: %v\n", score)
}

func Quiz(answerCh chan string) {
	var answer string
	fmt.Scanf("%s\n", &answer)
	answerCh <- answer
}

func parsedProblems(lines [][]string) []Problems {
	problems := make([]Problems, len(lines))
	for index, value := range lines {
		problems[index] = Problems{
			question: value[0],
			answer:   strings.TrimSpace(value[1]),
		}
	}
	return problems
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
