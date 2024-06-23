package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

// package main

// import (
// 	"fmt"
// 	"slices"
// )

// func main() {

// 	var s []string
// 	fmt.Println("uninit:", s, s == nil, len(s) == 0)

// 	s = make([]string, 3)
// 	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

// 	s[0] = "a"
// 	s[1] = "b"
// 	s[2] = "c"
// 	fmt.Println("set:", s)
// 	fmt.Println("get:", s[2])

// 	fmt.Println("len:", len(s))

// 	s = append(s, "d")
// 	s = append(s, "e", "f")
// 	fmt.Println("apd:", s)

// 	c := make([]string, len(s))
// 	copy(c, s)
// 	fmt.Println("cpy:", c)

// 	l := s[2:5]
// 	fmt.Println("sl1:", l)

// 	l = s[:5]
// 	fmt.Println("sl2:", l)

// 	l = s[2:]
// 	fmt.Println("sl3:", l)

// 	t := []string{"g", "h", "i"}
// 	fmt.Println("dcl:", t)

// 	t2 := []string{"g", "h", "i"}
// 	if slices.Equal(t, t2) {
// 		fmt.Println("t == t2")
// 	}

// 	twoD := make([][]int, 3)
// 	for i := 0; i < 3; i++ {
// 		innerLen := i + 1
// 		twoD[i] = make([]int, innerLen)
// 		for j := 0; j < innerLen; j++ {
// 			twoD[i][j] = i + j
// 		}
// 	}
// 	fmt.Println("2d: ", twoD)
// }

// // package main

// // import "fmt"

// // func main() {

// // 	if 7%2 == 0 {
// // 		fmt.Println("7 is even")
// // 	} else {
// // 		fmt.Println("7 is odd")
// // 	}

// // 	if 8%4 == 0 {
// // 		fmt.Println("8 is divisible by 4")
// // 	}

// // 	if 7%2 == 0 || 8%2 == 0 {
// // 		fmt.Println("either 8 or 7 are even")
// // 	}

// // 	if num := 9; num < 0 {
// // 		fmt.Println(num, "is negative")
// // 	} else if num < 10 {
// // 		fmt.Println(num, "has 1 digit")
// // 	} else {
// // 		fmt.Println(num, "has multiple digits")
// // 	}
// // }
