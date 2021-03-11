package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	for i := 0; i < 30; i++ {
		fmt.Println("Hello world")
	}
	endTime := time.Since(startTime)
	fmt.Printf("runtime %s", endTime)
}
