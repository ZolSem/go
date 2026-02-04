package main

import (
	"fmt"
	"sync"
)

var c int = 0

func main() {

	wg := &sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(10)
	go counter(wg)
	go counter(wg)
	go counter(wg)
	go counter(wg)
	go counter(wg)

	go counter(wg)
	go counter(wg)
	go counter(wg)
	go counter(wg)
	go counter(wg)

	fmt.Println(c)
}

func counter(wg *sync.WaitGroup) {
	defer wg.Done()
	for range 1000 {
		c++
	}
}
