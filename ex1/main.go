package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	fileFlag := flag.String("file", "./problems.csv", "specify a .csv file to open")
	timeFlag := flag.String("time", "30s", "enter a time limit to answer the quiz. eg 30s for 30 seconds")
	shuffleFlag := flag.Bool("shuffle", false, "shuffle the order the questions will be asked in")
	flag.Parse()

	file, err := os.Open(*fileFlag)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	problems, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	if *shuffleFlag {
		shuffle(problems)
	}

	correct := 0
	qAnswered := 0

	fmt.Println("press enter when ready!")
	fmt.Scanf("\n")

	quizDuration, _ := time.ParseDuration(*timeFlag)
	time.AfterFunc(quizDuration, func() {
		quizResults(len(problems), correct)
		os.Exit(0)
	})

	var answer int
	for _, p := range problems {
		fmt.Println(p[0])
		fmt.Scanf("%d\n", &answer)
		i, _ := strconv.Atoi(p[1])
		if answer == i {
			fmt.Println("Correct!")
			correct++
		} else {
			fmt.Println("Incorrect!")
		}
		qAnswered++
	}
	quizResults(len(problems), correct)
}

func quizResults(total, correct int) {
	fmt.Println("total questions:", total, "right answers:", correct)
}

func shuffle(problems [][]string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(problems), func(i, j int) {
		problems[i], problems[j] = problems[j], problems[i]
	})
}
