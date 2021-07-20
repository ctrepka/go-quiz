package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type question struct {
	question string
	answer   string
}

func parseQuizFromCSV(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("Unable to read input file "+path, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	questions, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+path, err)
	}

	return questions
}

func main() {
	questions := parseQuizFromCSV("quiz.csv")
	//fmt.Println(questions)
	var numCorrect int = 0

	for index, value := range questions {

		fmt.Println("Question " + strconv.Itoa(index) + ": " + value[0] + " = ?")

		reader := bufio.NewReader((os.Stdin))
		text, _ := reader.ReadString('\n')

		var isCorrect bool = strings.TrimSpace(text) == strings.TrimSpace(value[1])

		if isCorrect {
			numCorrect += 1
			fmt.Println("You answered '" + strings.TrimSpace(text) + "' which is correct! You've answered " + strconv.Itoa(int(numCorrect)) + " out of " + strconv.Itoa(len(questions)) + " questions correctly. \n\n")
		} else {
			fmt.Println("You answered '" + strings.TrimSpace(text) + "' which is incorrect :( You've answered " + strconv.Itoa(int(numCorrect)) + " out of " + strconv.Itoa(len(questions)) + " questions correctly. \n\n")
		}

		var percentCorrect float64 = float64(numCorrect) / float64(len(questions)) * 100

		fmt.Println("You have completed the quiz. You answered " + strconv.FormatFloat(percentCorrect, 'f', 1, 64) + " percent of questions correctly")

	}
}
