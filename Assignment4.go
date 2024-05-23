//Goroutine with Timeout Problem: Write a Go program that performs a task in a goroutine and waits for it to finish. However, if the task takes more than 3 seconds, the program should print a timeout message and exit.

package main

import (
    "fmt"
    "time"
)

func main() {
    done := make(chan bool)

    go performTask(done)

    select {
    case <-done:
        fmt.Println("Task completed successfully.")
    case <-time.After(3 * time.Second):
        fmt.Println("Task took too long  to complete.")
    }
}

func performTask(done chan<- bool) {
    time.Sleep(4 * time.Second)
    done <- true
}