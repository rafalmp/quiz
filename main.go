package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Deferred function do not run when os.Exit() is called - it exits immediately.
	// One method to ensure that defer calls will be invoked is to wrap os.Exit()
	// inside a defer function.
	exitCode := 0
	defer func() { os.Exit(exitCode) }()

	csvFileName := flag.String(
		"csv", "problems.csv", "A CSV file containing 'question,answer' in each line.",
	)
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Printf("Error opening CSV file: %s\n", err)
		exitCode = 1
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Printf("Error parsing CSV: %s\n", err)
		exitCode = 1
		return
	}

	fmt.Println(records)
}
