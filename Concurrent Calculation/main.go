package main

import (
	"fmt"
	"sync"
)

func sum(start, end int, wg *sync.WaitGroup, result chan int) {
	defer wg.Done()
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	result <- sum
}

func main() {
	var wg sync.WaitGroup
	result := make(chan int)

	wg.Add(2)
	go sum(1, 50, &wg, result)
	go sum(51, 100, &wg, result)

	go func() {
		wg.Wait()
		close(result)
	}()

	total := 0
	for partialSum := range result {
		total += partialSum
	}

	fmt.Println("Total sum:", total)
}
