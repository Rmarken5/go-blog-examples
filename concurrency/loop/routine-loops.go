package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func printLine(inx int, wg *sync.WaitGroup) {
	defer wg.Done()
	x := int(rand.Float32() * 10)
	time.Sleep(time.Duration(x) * time.Second)
	fmt.Printf("Index: %d\n", inx)

}

func main() {
	wg := sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go printLine(i, &wg)
	}
	wg.Wait()
	fmt.Printf("Done")
}
