/* Thoughts: 
	1) I could have parsed the file first, then started the quiz. The way I did it has two issues:
		a) If the user times out, I had to use a flag to tell the program to keep 
			parsing q/a pairs, but stop asking for responses. (To find the total # of
			questions, you have to get to the end of the file, even though the quiz
			is over.)
		b) No way to shuffle the questions.
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
	"io"
	"strings"
	"flag"
	"time"
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

	var timeOut bool = false
	score := 0
	total := 0

	fmt.Printf("Time limit: %d seconds. Please press <enter> to continue.\n",*timeLimit)
	_, _ = q.ReadString('\n')
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if len(record) != 2 {
			fmt.Println("Line: ", record, "does not appear to be a question-answer pair. Skipping")
		} else {
			record[1] = strings.TrimSpace(record[1])
			answerCh := make(chan string)
			total++
			if !timeOut {

				go func() {
					fmt.Print(record[0], " = ")
					text, _ := q.ReadString('\n')
					text = strings.Replace(text, "\n", "", -1)
					answerCh <- strings.TrimSpace(text)
				}()

				select {
				case <- timer.C:
					fmt.Println("\nTime's up!")
					timeOut = true
					break

				case answer := <- answerCh:
					if answer == record[1] {
						score++
					}
				}
			} 
		}
	}

	fmt.Printf("Score: %d out of %d correct. That's %.2f%%\n", score, total, 100.0*float64(score)/float64(total))
}
