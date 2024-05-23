//Goroutine with Channel Problem: Write a Go program that calculates the sum of numbers from 1 to N concurrently using goroutines and channels. The program should take the value of N as input from the user.

package main

import (
	"fmt"
	"sync"
)

func sum(start, end int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	ch <- sum
}

func main() {
	var n int
	fmt.Print("Enter a number (N): ")
	fmt.Scan(&n)

	ch := make(chan int, n)
	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go sum((i-1)*n/n+1, i*n/n, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	firstValue := <-ch // First value
	total := 0
	for val := range ch {
		total += val
	}
	lastValue := firstValue + total // Last value

	fmt.Printf("First value: %d\n", firstValue)
	fmt.Printf("Last value: %d\n", lastValue)
	fmt.Printf("Sum of numbers from 1 to %d is %d\n", n, total+firstValue+lastValue)
}