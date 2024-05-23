//Pipeline Pattern Problem: Implement a pipeline pattern in Go where one goroutine generates numbers, another squares them, and a third prints the squared numbers.

package main

import (
    "fmt"
)

func main() {

    numbers := make(chan int)
    squared := make(chan int)

    go generateNumbers(numbers)

    go squareNumbers(numbers, squared)

    printSquaredNumbers(squared)
}

func generateNumbers(out chan<- int) {
    defer close(out)
    for i := 1; i <= 5; i++ {
        out <- i
    }
}

func squareNumbers(in <-chan int, out chan<- int) {
    defer close(out)
    for num := range in {
        out <- num * num
    }
}
func printSquaredNumbers(in <-chan int) {
    for squaredNum := range in {
        fmt.Println(squaredNum)
    }
}
