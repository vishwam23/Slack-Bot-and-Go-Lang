package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	for scanner.Scan() {
		line := scanner.Text()

		sentences := strings.Split(line, ".")

		for _, sentence := range sentences {
			sentence = strings.TrimSpace(sentence)

			sentence = strings.ToLower(sentence)
			sentence = strings.Title(sentence)
			var cleanedSentence strings.Builder
			for _, char := range sentence {
				if unicode.IsLetter(char) || unicode.IsSpace(char) || unicode.IsNumber(char) || char == '.' {
					cleanedSentence.WriteRune(char)
				}
			}

			sentence = strings.Join(strings.Fields(cleanedSentence.String()), " ")

			fmt.Fprintln(outputFile, sentence)
		}
	}

	fmt.Println("Output file 'output.txt' has been created successfully.")
}
