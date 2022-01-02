package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		x := int(rand.Float64() * 10)
		time.Sleep(time.Duration(x) * time.Second)
		fmt.Println(x)
		fmt.Println("Function one done")

	}(&wg)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		x := int(rand.Float64() * 10)
		time.Sleep(time.Duration(x) * time.Second)
		fmt.Println(x)
		fmt.Println("Function two done")
	}(&wg)
	wg.Wait()
	fmt.Println("done")
}
