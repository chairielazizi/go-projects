package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv","problem.csv","a csv file in the format of question answer")
	timeLimit := flag.Int("limit",30,"the time limit for the quiz in seconds")
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
	//fmt.Println(lines)

	problems := parseLines(lines)
	//fmt.Println(problems)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	//<- timer.C // wait the message to enter the channel

	// print to the end user
	correct := 0
	for i, p:=range problems{
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func(){
			var answer string
			fmt.Scanf("%s\n",&answer)
			answerCh <- answer // send it to the answer channel
		}()

		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n",correct,len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				//fmt.Println("Correct!")
				correct++
			}

		//default:
		//	fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		//	var answer string
		//	fmt.Scanf("%s\n", &answer) //not suitable for strings and multiple words
		//	if answer == p.a {
		//		fmt.Println("Correct!")
		//		correct++
		//	}

		}
	}
	fmt.Printf("You scored %d out of %d.",correct,len(problems))
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines)) // slice
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]), // to ignore the space in the csv file
		}
	}
	return ret
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}