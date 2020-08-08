package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv","problem.csv","a csv file in the format of question answer")
	flag.Parse()
	_ = csvFilename

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Print("Failed to open the CSV file: %s", *csvFilename)
		os.Exit(1)
	}
	_ = file
}
