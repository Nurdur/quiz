package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"fmt"
	"io"
	"strings"
	"flag"
)


func main() {
	// Either open default or user-specified quiz file
	var problemsFile string
	flag.StringVar(&problemsFile,"file", "problems.csv", "name of csv file containing the set of questions and answers")
	flag.Parse()
	fmt.Println("Opening", problemsFile)
	f, _ := os.Open(problemsFile)

	// Create a reader
	r := csv.NewReader(bufio.NewReader(f))
	q := bufio.NewReader(os.Stdin)
	score := 0
	total := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if len(record) != 2 {
			fmt.Println("Line: ", record, "does not appear to be a question-answer pair. Skipping")
		} else {
			fmt.Print(record[0], " = ")
			text, _ := q.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			total++
			if text == record[1] {
				score++
			} 
		}
	}
	fmt.Printf("Score: %d out of %d correct. That's %.2f%%\n", score, total, 100.0*float64(score)/float64(total))
}
