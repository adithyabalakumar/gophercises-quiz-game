package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Problem struct {
	Question string
	Answer   string
}

func printHelp() {
	fmt.Println("Usage: quiz [options]")
	fmt.Println("Options:")
	fmt.Println("  -h,		Show this help message")
	fmt.Println("  -file,	CSV file containing quiz questions and answers (default: problems.csv)")
}

func readCSV(fileNamePtr *string) ([]Problem, error) {
	// Open the CSV file
	// Read the contents of the CSV file
	// Parse the questions and answers
	// Store them in a suitable data structure (e.g., a slice of structs)
	data, err := os.ReadFile(*fileNamePtr)
	if err != nil {
		return nil, err
	}
	// Process the CSV data here (using encoding/csv package)
	r := csv.NewReader(strings.NewReader(string(data)))
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	questions := make([]Problem, 0)
	for _, record := range records {
		if len(record) == 2 {
			questions = append(questions, Problem{
				Question: strings.TrimSpace(record[0]),
				Answer:   strings.TrimSpace(record[1]),
			})
		}
	}
	return questions, nil
}

func startQuiz(questions []Problem) {
	// Loop through the questions and present them to the user
	// Collect the user's answers and keep track of the score
	score := 0
	for i, q := range questions {
		fmt.Printf("Question #%d: %s\nYour answer: ", i+1, q.Question)
		var userAnswer string
		fmt.Scanln(&userAnswer)
		if strings.TrimSpace(userAnswer) == q.Answer {
			score++
		}
	}
	fmt.Printf("Quiz completed! Your score is: %d/%d\n", score, len(questions))
}

func main() {
	// Get the command-line arguments
	helpPtr := flag.Bool("h", false, "Show this help message")
	fileNamePtr := flag.String("file", "problems.csv", "CSV file containing quiz questions and answers")
	flag.Parse()

	if *helpPtr {
		printHelp()
		return
	}

	fmt.Println("Using file:", *fileNamePtr)
	questions, err := readCSV(fileNamePtr)
	if err != nil {
		fmt.Println("Failed to read questions from CSV file:", err)
		return
	}

	startQuiz(questions)
}
