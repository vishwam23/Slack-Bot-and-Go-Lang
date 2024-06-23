package main

import (
	"fmt"
	"time"
)

func main() {
	go CLock()

	select {}
}

func CLock() {
	for {
		currentTime := time.Now()
		hour, minute, second := currentTime.Hour(), currentTime.Minute(), currentTime.Second()
		fmt.Printf("\rtime : %02d:%02d:%02d", hour, minute, second)
		time.Sleep(1 * time.Second)
	}
}
