package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Emails    []string
	Addresses map[string]string
}

func PrintPersonDetails(p Person) {
	fmt.Println("First Name:", p.FirstName)
	fmt.Println("Last Name:", p.LastName)
	fmt.Println("Age:", p.Age)

	fmt.Println("Emails:")
	for _, email := range p.Emails {
		fmt.Println("- ", email)
	}

	fmt.Println("Addresses:")
	for label, address := range p.Addresses {
		fmt.Printf("- %s: %s\n", label, address)
	}
}

func main() {
	person := Person{}

	person.FirstName = "vishwam"
	person.LastName = "arora"
	person.Age = 30
	person.Emails = []string{"vishwam23@example.com", "vishwamarora@gmail.com"}
	person.Addresses = map[string]string{
		"Home":     "38 Radaur",
		"Work":     "Bonovalley,Delhi",
		"Vacation": "Manali,HP",
	}

	PrintPersonDetails(person)
}
