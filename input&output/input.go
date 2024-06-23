package main

import (
	"fmt"
	"sync"
)

func takeInput(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	inputs := []string{"one", "two", "three", "four", "five"}

	for _, input := range inputs {
		ch <- input
	}
	close(ch)
}

func printInput(ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for input := range ch {
		fmt.Println("Received input:", input)
	}
}

func main() {
	var wg sync.WaitGroup

	inputChannel := make(chan string, 3)

	wg.Add(2)

	go takeInput(inputChannel, &wg)
	go printInput(inputChannel, &wg)

	wg.Wait()
}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func takeInput(ch chan<- string) {
// 	var input string
// 	fmt.Print("Enter some input: ")
// 	fmt.Scanln(&input)
// 	ch <- input
// 	close(ch)
// }

// func main() {
// 	inputChannel := make(chan string)

// 	go takeInput(inputChannel)

// 	time.Sleep(2 * time.Second)

// 	input := <-inputChannel
// 	fmt.Println("Received input:", input)
// }
