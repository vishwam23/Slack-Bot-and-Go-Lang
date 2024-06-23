package main

import "fmt"

func myfunc(ch chan int) {

	fmt.Println(2304 + <-ch)
}
func main() {
	fmt.Println("start Main method")
	ch := make(chan int)

	go myfunc(ch)
	ch <- 2002
	fmt.Println("End Main method")
}
