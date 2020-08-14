package main

import (
	"encoding/csv"
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
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	//_ = file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	fmt.Println(lines)

	problems := parseLines(lines)
	fmt.Println(problems)

	// print to the end user
	for i, p:=range problems{
		fmt.Printf("Problem #%d: %s = \n",i+1,p.q)
		var answer string
		fmt.Scanf("%s\n",&answer) //not suitable for strings and multiple words
		if answer == p.a{
			fmt.Println("Correct!")
		}
	}
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}