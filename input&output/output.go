package main

import (
	"fmt"
	"time"
)

func showOutput(ch <-chan string) {
	time.Sleep(3 * time.Second)

	output := <-ch
	fmt.Println("Processed input:", output)
}

func main() {
	outputChannel := make(chan string)

	go showOutput(outputChannel)

	var input string
	fmt.Print("Enter some input: ")
	fmt.Scanln(&input)
	outputChannel <- input

	time.Sleep(4 * time.Second)
}
