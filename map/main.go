package main

import "fmt"

func main() {

	Animal := map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}
	fmt.Println("Original map: ", Animal)

	value_1 := Animal[90]
	value_2 := Animal[93]
	fmt.Println("Value of key[90]: ", value_1)
	fmt.Println("Value of key[93]: ", value_2)
}
