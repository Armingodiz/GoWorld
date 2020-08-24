package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var questions, rightAnswers int
var timeLimit *int
var quiz = make(map[string]string)

func init() {
	csvFile, err := os.Open("problems.csv")
	reader := csv.NewReader(csvFile)
	checkErr(err)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		questions++
		quiz[record[0]] = record[1]
	}
	timeLimit = flag.Int("time", 30, "time limit for quiz")
}
func main() {
	flag.Parse()
	timer1 := time.NewTimer(3014 * time.Second)

	// Read each record from csv
	go ask()
	for {
		select {
		case <-timer1.C:
			fmt.Println("number of Question ASKED  : ", questions)
			fmt.Println("Right Answers : ", rightAnswers)
			return
		default:
		}
	}
}
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func ask() {
	for q, _ := range quiz {
		fmt.Println("Question:", q)
		var answer int
		fmt.Scan(&answer)
		rigthAnswer, err2 := strconv.Atoi(quiz[q])
		checkErr(err2)
		if answer == rigthAnswer {
			rightAnswers++
		}
	}
	fmt.Println("number of Question : ", questions)
	fmt.Println("Right Answers : ", rightAnswers)
}
