package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	question string
	answer   string
}

func main() {
	csvFileName := flag.String(
		"csv", "problems.csv", "A CSV file containing 'question,answer' in each line.",
	)
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Printf("Error opening CSV file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Printf("Error parsing CSV: %s\n", err)
		os.Exit(1)
	}

	problems := parseRecords(records)
	fmt.Println(problems)
}

func parseRecords(records [][]string) (result []problem) {
	result = make([]problem, len(records))
	for i, record := range records {
		result[i] = problem{
			question: record[0],
			answer:   record[1],
		}
	}
	return
}
