package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"fmt"
	"io"
	"strings"
)


func main() {
	// Open the file "problems.csv"
	// TODO Accept a flag to load a user-specified file instead
	f, _ := os.Open("./problems.csv")

	// Create a reader
	r := csv.NewReader(bufio.NewReader(f))
	q := bufio.NewReader(os.Stdin)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if len(record) != 2 {
			fmt.Println("Line: ", record, "does not appear to be a question-answer pair. Skipping")
		} else {
			fmt.Println("Question: ", record[0])
			fmt.Print("-> ")
			text, _ := q.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			fmt.Println("Response: ", text)
			fmt.Println("Answer: ", record[1])
		}
	}
}
