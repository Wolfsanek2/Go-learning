package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 100)
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		ch <- i
	}
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		for v := range ch {
			fmt.Println("Горутина 1 получила", v)
		}
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		for v := range ch {
			fmt.Println("Горутина 2 получила", v)
		}
		wg.Done()
	}(wg)
	close(ch)
	wg.Wait()
}
