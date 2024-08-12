// counting how many times random numbers are between 0 and 9 using sync.Map
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
)

func generateRandomNumber(max int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()), nil
}

func updateCount(countMap *sync.Map, key int) {
	count, _ := countMap.LoadOrStore(key, 0)
	countMap.Store(key, count.(int)+1)
}

func printCounts(countMap *sync.Map) {
	countMap.Range(func(key, value any) bool {
		fmt.Printf("Number %d: Count %d\n", key, value)
		return true
	})
}

func main() {
	var countMap sync.Map
	numGoroutines := 5
	var wg sync.WaitGroup
	generateAndCount := func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			randomNumber, err := generateRandomNumber(10)
			if err != nil {
				fmt.Println("Error generating number:", err)
				return
			}
			updateCount(&countMap, randomNumber)
		}
	}
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go generateAndCount()
	}
	wg.Wait()
	printCounts(&countMap)
}
