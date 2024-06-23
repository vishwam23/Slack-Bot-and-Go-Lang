package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	City  string `json:"city"`
	Email string `json:"email"`
}

func main() {
	jsonData := `{"name":"Vishwam","age":21,"city":"Ynr","email":"vishwam23@gmail.com"}`

	var personStruct Person
	err := json.Unmarshal([]byte(jsonData), &personStruct)
	if err != nil {
		fmt.Println("Error unmarshalling ", err)
		return
	}

	fmt.Println("Unmarshalled JSON into struct:")
	fmt.Println("Name:", personStruct.Name)
	fmt.Println("Age:", personStruct.Age)
	fmt.Println("City:", personStruct.City)
	fmt.Println("Email:", personStruct.Email)
	fmt.Println()

	var personInterface interface{}
	err = json.Unmarshal([]byte(jsonData), &personInterface)
	if err != nil {
		fmt.Println("Error unmarshalling JSON into interface:", err)
		return
	}

	fmt.Println("Unmarshalled JSON into interface:")
	personMap, ok := personInterface.(map[string]interface{})
	if !ok {
		fmt.Println("Error asserting interface to map[string]interface{}")
		return
	}

	for key, value := range personMap {
		fmt.Printf("%s: %v\n", key, value)
	}
}
