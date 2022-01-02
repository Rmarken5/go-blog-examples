package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printLine(inx int) {
	x := int(rand.Float32() * 10)
	time.Sleep(time.Duration(x) * time.Second)
	fmt.Printf("Index: %d\n", inx)
}

func main() {
	var s string
	for i := 1; i <= 10; i++ {
		go printLine(i)
	}
	fmt.Scanln(&s)
	fmt.Printf("Done")
}
