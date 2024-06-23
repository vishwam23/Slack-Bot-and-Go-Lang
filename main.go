package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello brother")
	var Number int = 2342
	fmt.Println(Number)
	fmt.Printf("type is %T \n", Number)
	no := 123
	fmt.Printf("type is %T %d \n", no, no)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter no. ")
	input, _ := reader.ReadString('\n')
	fmt.Println("your no. is ", input)
	fmt.Println("enter your name ")
	input, _ = reader.ReadString('\n')
	fmt.Println("your name is", input)

	fmt.Println("enter no.")
	input, _ = reader.ReadString('\n')
	add, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("add 1 to no", add+1)
	}

}
