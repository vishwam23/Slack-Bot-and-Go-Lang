package main

import "fmt"

type CustomStruct struct {
	Field1 int
	Field2 string
}

func main() {
	var mixedSlice []interface{}

	mixedSlice = append(mixedSlice, 42)                             // int
	mixedSlice = append(mixedSlice, 3.14)                           // float64
	mixedSlice = append(mixedSlice, "Hello, World!")                // string
	mixedSlice = append(mixedSlice, map[string]int{"a": 1, "b": 2}) // map
	mixedSlice = append(mixedSlice, CustomStruct{10, "Example"})    // struct

	for _, value := range mixedSlice {
		switch v := value.(type) {
		case int:
			fmt.Println("int, Value:", v)
		case float64:
			fmt.Println("float64, Value:", v)
		case string:
			fmt.Println("string, Value:", v)
		case map[string]int:
			fmt.Println("map[string]int, Value:", v)
		case CustomStruct:
			fmt.Println("CustomStruct, Value:", v)
		default:
			fmt.Println("unknown, Value:", v)
		}
	}
}
