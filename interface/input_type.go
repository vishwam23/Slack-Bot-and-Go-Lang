package main

import "fmt"

func PrintParameter[T int | float64 | string](param T) {
	fmt.Println("value:", param)
}

func main() {
	PrintParameter(2332)
	PrintParameter(334.4554)
	PrintParameter("Vishwam")
}
