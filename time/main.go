package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("time is", t.Format("01-02-2006 Monday 15:04:05"))

}
