package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type quiz struct {
	challenge, response string
}

func main() {

	filename := flag.String("csv", "problems.csv", "CSV File consist of quizzes.")
	timeLimit := flag.Int("limit", 30, "Time limit for the quiz.")
	flag.Parse()

	qs, err := readCSV(*filename)
	if err != nil {
		log.Fatal(err)
	}
	startQuiz(qs, *timeLimit)
}

func startQuiz(qs []quiz, timeLimit int) {

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	score := 0
	for i, q := range qs {
		fmt.Printf("Problem #%d: %s = ", i+1, q.challenge)

		ansC := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			ansC <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d", score, len(qs))
			return
		case ans := <-ansC:
			if ans == q.response {
				score++
			}
		}

	}
}

func readCSV(filename string) ([]quiz, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %s", filename)
	}
	defer file.Close()

	r := csv.NewReader(file)

	out := []quiz{}

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("error while reading CSV record: %v", record)
		}

		if len(record) != 2 {
			return nil, fmt.Errorf("unexpected number of fields for record: %v", record)
		}

		out = append(out, quiz{record[0], record[1]})
	}
	return out, nil
}
