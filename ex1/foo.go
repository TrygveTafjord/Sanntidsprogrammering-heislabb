// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
)

func incrementing(channel chan bool, quit chan bool) {
	//TODO: increment i 1000000 times
	for j := 0; j < 1000000; j++ {
		channel <- true
	}
	quit <- true
}

func decrementing(channel chan bool, quit chan bool) {
	//TODO: decrement i 1000000 times
	for j := 0; j < 1000000; j++ {
		channel <- true
	}
	quit <- true
}

func server(ch1 chan bool, ch2 chan bool, read chan int) {
	var i = 0
	for {
		select {
		case <-ch1:
			i++
		case <-ch2:
			i--
		case read <- i:
		}
	}
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1? //sets one CPU
	runtime.GOMAXPROCS(2)

	increment := make(chan bool)
	decrement := make(chan bool)
	quit := make(chan bool)
	read := make(chan int)

	// TODO: Spawn both functions as goroutines
	go incrementing(increment, quit)
	go decrementing(decrement, quit)
	go server(increment, decrement, read)

	<-quit
	<-quit

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.

	Println("The magic number is:", <-read)
}
