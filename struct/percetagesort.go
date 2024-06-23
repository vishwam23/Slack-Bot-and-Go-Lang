// package main

// import (
// 	"fmt"
// 	"sort"
// )

// type Student struct {
// 	Name       string
// 	Percentage float64
// }

// type ByPercentage []Student

// func (s ByPercentage) Len() int           { return len(s) }
// func (s ByPercentage) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
// func (s ByPercentage) Less(i, j int) bool { return s[i].Percentage < s[j].Percentage }

// func main() {
// 	students := []Student{
// 		{Name: "A", Percentage: 78.5},
// 		{Name: "B", Percentage: 82.3},
// 		{Name: "C", Percentage: 76.8},
// 		{Name: "D", Percentage: 91.2},
// 		{Name: "E", Percentage: 85.0},
// 	}

// 	sort.Sort(ByPercentage(students))
// 	fmt.Println("sorted in ascending order of percentage:")
// 	for _, student := range students {
// 		fmt.Printf("%s: %.2f%%\n", student.Name, student.Percentage)
// 	}

//		sort.Sort(sort.Reverse(ByPercentage(students)))
//		fmt.Println("\nS sorted in descending order of percentage:")
//		for _, student := range students {
//			fmt.Printf("%s: %.2f%%\n", student.Name, student.Percentage)
//		}
//	}
package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name       string
	Percentage float64
}

func main() {
	students := []Student{
		{Name: "A", Percentage: 78.5},
		{Name: "B", Percentage: 82.3},
		{Name: "C", Percentage: 76.8},
		{Name: "D", Percentage: 91.2},
		{Name: "E", Percentage: 85.0},
	}

	sort.Slice(students, func(i, j int) bool {
		return students[i].Percentage < students[j].Percentage
	})
	printStudents(students)

	sort.Slice(students, func(i, j int) bool {
		return students[i].Percentage > students[j].Percentage
	})
	fmt.Println("\nStudents sorted in descending order of percentage:")
	printStudents(students)
}

func printStudents(students []Student) {
	for _, student := range students {
		fmt.Printf("%s: %.2f%%\n", student.Name, student.Percentage)
	}
}
