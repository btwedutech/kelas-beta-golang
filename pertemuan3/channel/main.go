package main

import (
	"fmt"
	"sync"
)

func main() {
	var ch = make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(2)
	go receiveMessage(ch, 1, &wg)
	go receiveMessage(ch, 2, &wg)

	for i := 0; i < 40; i++ {
		ch <- i
	}
	close(ch)

	wg.Wait()

}

func receiveMessage(ch chan int, name int, wg *sync.WaitGroup) {
	for msg := range ch {
		fmt.Println("Data", msg, "Channel", name)
	}

	wg.Done()
}
