package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type quiz struct {
	challenge, response string
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println(`USAGE: go run main.go <CSV_FILE>`)
		fmt.Println(os.Args)
		return
	}

	qs, err := readCSV(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	score := 0

	for i, q := range qs {
		fmt.Printf("Problem #%d: %s = \n", i+1, q.challenge)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if q.response == answer {
			score++
		}
	}
	fmt.Printf("You scored %d out of %d", score, len(qs))
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
