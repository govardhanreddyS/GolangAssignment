//Mutex for Synchronization Problem: Write a Go program that uses mutexes to synchronize access to a shared variable. Multiple goroutines should increment the variable concurrently, and the final value should be printed.

package main

import (
    "fmt"
    "sync"
    "time"
)

var (
    counter int
    mutex   sync.Mutex
    wg      sync.WaitGroup
)

func main() {
    
    numGoroutines := 5
    
    wg.Add(numGoroutines)
    
    for i := 0; i < numGoroutines; i++ {
        go incrementCounter(i)
    }
    
    wg.Wait()
    
    fmt.Println("Final Counter:", counter)
}

func incrementCounter(id int) {
    defer wg.Done()
    
    time.Sleep(time.Duration(id) * time.Millisecond)
    
    mutex.Lock()
    counter++
    mutex.Unlock()
}
