package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	channel := make(chan int)
	
	wg.Add(2)
	go receiveData(channel, 1, &wg)
	go receiveData(channel, 2, &wg)

	for i := 0; i < 10; i++ {
		channel <- i
	}

	close(channel)

	wg.Wait()
}

func receiveData(channel chan int, number int, wg *sync.WaitGroup) {
	for data := range channel {
		fmt.Println("Data received ", data, "From Goroutine ", number)
	}
	wg.Done()
}
