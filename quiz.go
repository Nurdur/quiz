/* Thoughts: 
	1) Shuffle the questions?
	2) I'm still not super comfortable with:
		a) go routines
		b) channels
		c) bufio
	3) Cool way to learn about:
		a) file and console io
		b) flags
		c) fmt
		d) the simple basics of go
		e) git branching
	I'm glad I did this one. Did a little at a time over two days.
*/

package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"fmt"
	"strings"
	"flag"
	"time"
	"log"
)


func main() {

	problemsFile := flag.String("file", "problems.csv", "name of csv file containing the set of questions and answers")
	timeLimit := flag.Int("time",30,"time limit")
	flag.Parse()

	fmt.Println("Opening", *problemsFile)
	f, _ := os.Open(*problemsFile)

	// Create a reader
	r := csv.NewReader(bufio.NewReader(f))
	q := bufio.NewReader(os.Stdin)

	score := 0

	fmt.Printf("Time limit: %d seconds. Please press <enter> to continue.\n",*timeLimit)
	_, _ = q.ReadString('\n')
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

//	Print for debugging
//	fmt.Print(records)


mainLoop:
	for _, record := range records  {
		if len(record) != 2 {
			fmt.Println("Line: ", record, "does not appear to be a question-answer pair. Skipping")
		} else {
			record[1] = strings.TrimSpace(record[1])
			answerCh := make(chan string)

			go func() {
				fmt.Print(record[0], " = ")
				text, _ := q.ReadString('\n')
				text = strings.Replace(text, "\n", "", -1)
				answerCh <- strings.TrimSpace(text)
			}()

			select {
			case <- timer.C:
				fmt.Println("\nTime's up!")
				break mainLoop

			case answer := <- answerCh:
				if answer == record[1] {
					score++
				}
			}
		}
	}

	fmt.Printf("Score: %d out of %d correct. That's %.2f%%\n", score, len(records), 100.0*float64(score)/float64(len(records)))
}
