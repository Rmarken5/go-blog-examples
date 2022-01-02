package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleepThread(inx int, out chan<- string) {
	x := int(rand.Float32() * 10)
	time.Sleep(time.Duration(x) * time.Second)
	out <- fmt.Sprintf("Index: %d has slept for %d seconds\n", inx, x)
}

func main() {
	outputChannel := make(chan string)

	for i := 1; i <= 10; i++ {
		go sleepThread(i, outputChannel)
	}

	for i := 1; i <= 10; i++ {
		fmt.Print(<-outputChannel)
	}

}
