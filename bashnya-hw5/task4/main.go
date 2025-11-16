package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	inputCh := make(chan int)
	outputCh := make(chan int)

	go func() {
		defer close(inputCh)
		for _, num := range numbers {
			inputCh <- num
		}
	}()

	go func() {
		defer close(outputCh)
		for num := range inputCh {
			result := num * 2
			outputCh <- result
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("Результаты:\n")
		for result := range outputCh {
			fmt.Printf("%d\n", result)
		}
	}()

	wg.Wait()
}
