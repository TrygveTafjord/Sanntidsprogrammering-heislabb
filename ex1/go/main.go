package main

import (
	"fmt"
	"time"
)

func producer(ch chan int, quit chan bool) {

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[producer]: pushing %d\n", i)
		ch <- i
	}
	quit <- true

}

func consumer(ch chan int,  quit chan bool) {

	time.Sleep(1 * time.Second)
	for {
		i := 0
		i = <-ch
		fmt.Printf("[consumer]: %d\n", i)
		time.Sleep(50 * time.Millisecond)

	}
}

func main() {

	// TODO: make a bounded buffer
	ch := make(chan int, 5)
    quit := make(chan bool)



	go consumer(ch,quit)
	go producer(ch, quit)

    <-quit

    fmt.Printf("Finito")

}