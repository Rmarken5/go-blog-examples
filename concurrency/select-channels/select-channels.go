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
	sleepChan := make(chan string)
	helloChan := make(chan string)

	for i := 1; i <= 10; i++ {
		go sleepThread(i, sleepChan)
		go func(inx int, out chan<- string) {
			x := int(rand.Float32() * 10)
			time.Sleep(time.Duration(x) * time.Second)
			out <- fmt.Sprintf("Index: %d says hello after %d seconds\n", inx, x)
		}(i, sleepChan)
	}
	go func() {
		for {
			select {
			case out := <-sleepChan:
				fmt.Print(out)
			case out := <-helloChan:
				fmt.Print(out)
			}
		}
	}()
	var s string
	fmt.Scanln(&s)
}
