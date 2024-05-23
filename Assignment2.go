//Producer-Consumer with Channel Problem: Implement the producer-consumer problem using goroutines and channels. The producer should generate numbers from 1 to 100 and send them to a channel, and the consumer should print those numbers.


package main

import (
	"fmt"
)

func producer(Ch chan<- int) {
	for i := 1; i <= 100; i++ {
		Ch <- i
	}
	close(Ch)
}

func consumer(Ch <-chan int) {
	for num := range Ch {
		fmt.Println("Numbers:", num)
	}
}

func main() {
	Ch := make(chan int)
	go producer(Ch)
	consumer(Ch)
}