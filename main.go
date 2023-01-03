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
	wd, err := os.Getwd()
	if err != nil {
		Exit("wd err")
	}
	defautlFileName := strings.Join([]string{wd, "problems.csv"}, "/")
	csvFileName := flag.String("csv", defautlFileName, "a name of the csv file formated 'question,anser'")
	file, err := os.Open(*csvFileName)

	if err != nil {
		Exit(fmt.Sprintf("failed to open file name: %v. ", *csvFileName))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		Exit("faild to parse file")
	}
	problems := ParsedProblems(lines)
	score := 0
	timer := time.NewTimer(time.Duration(10) * time.Second)

	for index, problem := range problems {
		fmt.Printf("Problem #%d: %s= ", index+1, problem.question)
		answerCh := make(chan string)
		var answer string
		go Quiz(answerCh)
		select {
		case <-timer.C:
			fmt.Printf("\nTime Over \nToal Score: %v\n", score)
			return
		case answer = <-answerCh:
			if answer != problem.answer {
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

func ParsedProblems(lines [][]string) []Problems {
	problems := make([]Problems, len(lines))
	for index, value := range lines {
		problems[index] = Problems{
			question: value[0],
			answer:   strings.TrimSpace(value[1]),
		}
	}
	return problems
}

func Exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
