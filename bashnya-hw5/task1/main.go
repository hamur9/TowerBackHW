package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	results := make(chan int)

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			results <- x * x
		}(num)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	sum := 0
	for square := range results {
		sum += square
	}

	fmt.Printf("Сумма квадратов: %d\n", sum)
}
