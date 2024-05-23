package main







import "fmt"

func main(){

    greeting := func(){

        FormData.Println("Hello world)
    }

    greeting();
}



func SharedWorker(){
    FormData.Println("Hello World")
}

func makeGreeter()function()string{
return func() string{
    return "Hello World"
}

}



func main() {
	fmt.Println(greet("jane","Doe",35))
	f,s:=("first","second",4)
	fmt.Print(f,s,a)
}

func greet(fname,lname,age)

........closures.........

package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}


import "fmt"

func visit(numbers []int, callback func(int)) {
	for i, n := range numbers {
		fmt.Print(i, n)
		callback(n)
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}

package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	var xs []int
	for _,n := range numbers {
		if callback(n) {
			xs = append(xs,n)
		}
	}
	return xs
}
func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs)
}



func filter(numbers []int, callback func(int) bool) []int {
	var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}
func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs)
}

// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	for n := range c {
		fmt.Print(n)
	}
}


func main() {
	msgQueue := make(chan string, 3)
	msgQueue <- "One"
	msgQueue <- "Two"
	msgQueue <- "Three"

	fmt.Print(<-msgQueue)
	fmt.Print(<-msgQueue)
	fmt.Print(<-msgQueue)

}



// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()
	<-done
	<-done
}


// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("sender:Sending Data : ", i)
			ch <- i
			fmt.Println("Sender : Data Sent : ", i)
			time.Sleep(1 * time.Second)
		}
		close(ch)
	}()

	go func() {
		fmt.Println("Receiver: Waiting for Data...")
		for data := range ch {
			fmt.Println("Receiver : Receiver data", data)
			time.Sleep(3 * time.Second)
		}
		fmt.Println("receiver is closed")
	}()
	time.Sleep(20 * time.Second)
}



package main

import "fmt"

func main() {
	c := incrementor()
	cSum := puller(c)
	for n := range cSum {
		fmt.Println(n)
	}
}

func incrementor() chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}


package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }
func main() {
	studyGroup := people{"zeno", "jhon", "Ai", "AIML", "jenny", "z"}
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
}



package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name   string
	Age    int
	Salary int
}

func (p person) String() string {
	return fmt.Sprintf("YAYAYA %s: %d", p.Name, p.Salary)
}

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Salary < a[j].Salary }
func main() {
	people := []person{
		{"Bob", 31, 2000},
		{"John", 42, 5000},
		{"Michael", 17, 5000},
		{"Jenny", 26, 6000},
	}

	fmt.Println(people[0])
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

}



package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

type rectangle struct {
	length float64
	width  float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.length * r.width

}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	r := rectangle{5, 10}
	info(r)
}



package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	pingCh := make(chan string)
	pongCh := make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 26; i++ {
			letter := string(rune('A' + i))
			fmt.Println("Ping", i)
			pongCh <- letter
			<-pingCh
		}
	}()

	go func() {
		wg.Add(1)

		defer wg.Done()
		for i := 0; i < 26; i++ {
			letter := <-pongCh
			fmt.Println("Pong", letter)
			pingCh <- string(rune(letter[0] + 1))
		}
	}()
	wg.Wait()
	close(pingCh)
	close(pongCh)
}




package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	pingCh := make(chan string)
	pongCh := make(chan string)

	wg.Add(1) //One WaitGroup for the first goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < 26; i++ { //Ping
			letter := string(rune('A' + i)) // Convert integer to corresponding letter
			// fmt.Println("Ping",i)
			fmt.Println("Ping ", string(letter[0]), i+1)
			pongCh <- letter
			<-pingCh
		}
	}()
	go func() { //Pong
		wg.Add(1) //One WaitGroup for the second goroutine
		defer wg.Done()
		for i := 0; i < 26; i++ {
			letter := <-pongCh
			// fmt.Println("Pong",letter)
			fmt.Println("Pong ", letter, i+1)
			pingCh <- string(rune(letter[0] + 1)) //Send next letter in sequence
		}
	}()
	wg.Wait() //Wait for the first goroutine to finish
	close(pingCh)
	close(pongCh)
}



package main

import (
	"fmt"
	"time"
)

func goroutineA(ch chan bool) {
	for {
		select {
		case <-ch:
			fmt.Println("Goroutine A is triggered")
			case <- quit:
			fmt
			
		}
	}
}

func goroutineB(ch chan bool) {
	for {
		select {
		case <-ch:
			fmt.Println("Goroutine B is triggered")
		}
	}
}

func main() {
	trigger := make(chan bool)

	go goroutineA(trigger)
	go goroutineB(trigger)

	go func() {
		for {
			time.Sleep(2 * time.Second)
			trigger <- true
			time.Sleep(3 * time.Second)
			trigger <- true
		}
	}()
	select {}
	time.Sleep(15*time.Second)
	close(quit)
	
}


package main

import (
	"fmt"
	"sync"
	"time"
)

func routine(id int, wg *sync.WaitGroup, ch chan string, pauseCh <-chan bool) {
	defer wg.Done()
	for {
		select {
		case <-pauseCh:
			fmt.Printf("Routine %d paused\n", id)
			<-pauseCh // Wait for unpause signal
			fmt.Printf("Routine %d resumed\n", id)
		default:
			fmt.Printf("Routine %d running\n", id)
			// Simulate some work
			time.Sleep(time.Second * 20)
			ch <- fmt.Sprintf("Routine %d completed", id)
		}
	}
}

func main() {
	numRoutines := 3
	ch := make(chan string, numRoutines)
	pauseCh := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(numRoutines)

	// Start routines
	for i := 1; i <= numRoutines; i++ {
		go routine(i, &wg, ch, pauseCh)
	}

	// Control loop
	for {
		// Round-robin scheduling
		select {
		case result := <-ch:
			fmt.Println(result)
		case <-time.After(5 * time.Second):
			// Pause all routines for 2 seconds
			fmt.Println("Pausing all routines...")
			for i := 0; i < numRoutines; i++ {
				pauseCh <- true
			}
			time.Sleep(3 * time.Second)
			// Resume all routines
			fmt.Println("Resuming all routines...")
			for i := 0; i < numRoutines; i++ {
				pauseCh <- false
			}
		}
	}

	// Wait for routines to finish (which will never happen in this example)
	wg.Wait()
}


package main

import (
	"fmt"
	"math/rand"
	//"time"
)

// gen generates random numbers and sends them to the returned channel
func gen() <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			out <- rand.Intn(100)
		}
	}()
	return out
}

// fanIn combines multiple input channels into a single channel
func fanIn(inputs ...<-chan int) <-chan int {
	out := make(chan int)
	for _, in := range inputs {
		go func(ch <-chan int) {
			for n := range ch {
				out <- n
			}
		}(in)
	}
	return out
}

func main() {
	// Create input channels
	ch1 := gen()
	ch2 := gen()
	ch3 := gen()

	// Combine input channels into a single channel
	fanInCh := fanIn(ch1, ch2, ch3)

	// Read values from the combined channel
	for i := 0; i < 10; i++ {
		fmt.Println(<-fanInCh)
	}
}


package main

import (
	"fmt"
	"math/rand"
	//"time"
)

// gen generates random numbers and sends them to the returned channel
func gen() <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			out <- rand.Intn(100)
		}
	}()
	return out
}

// fanIn combines multiple input channels into a single channel
func fanIn(inputs ...<-chan int) <-chan int {
	out := make(chan int)
	for _, in := range inputs {
		go func(ch <-chan int) {
			for n := range ch {
				out <- n
			}
		}(in)
	}
	return out
}

func main() {
	// Create input channels
	ch1 := gen()
	ch2 := gen()
	ch3 := gen()

	// Combine input channels into a single channel
	fanInCh := fanIn(ch1, ch2, ch3)

	// Read values from the combined channel
	for i := 0; i < 10; i++ {
		fmt.Println(<-fanInCh)
	}
}
