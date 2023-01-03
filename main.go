package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file name in the format of 'questions,answer'")
	limit := flag.Int("limit", 3, "duratioion of quiz game")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("failed to open CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse provided file")
	}
	problems := ParseLines(lines)
	correct := 0
	timer := time.NewTimer(time.Duration(*limit) * time.Second)

	for i, p := range problems {

		fmt.Printf("problem #%d: %s=\n", i+1, p.q)
		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s \n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\n time over \n total score: %d.\n", correct)
			return
		case answer := <-answerCh:

			if answer == p.a {
				correct++
				fmt.Println("correct")
			} else {
				fmt.Println("incorrect")
			}

		}

	}
	fmt.Printf("total score: %d.\n", correct)

}

func ParseLines(lines [][]string) []Problems {
	ret := make([]Problems, len(lines))
	for index, values := range lines {
		ret[index] = Problems{
			q: values[0],
			a: values[1],
		}
	}
	return ret
}

type Problems struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
