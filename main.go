package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file name in the format of 'questions,answer'")
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
	parsedLines := ParseLines(lines)
	correct := 0
	for i, p := range parsedLines {
		fmt.Printf("problem #%d: %s=\n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
			fmt.Println("correct")
		} else {
			fmt.Println("incorrect")
		}
	}
	fmt.Printf("You scored %d.\n", correct)

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
