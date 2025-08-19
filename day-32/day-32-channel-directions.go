package main

import "fmt"

func sendData(ch chan<- int) { // send-only channel
	ch <- 100
}

func receiveData(ch <-chan int) { // receive-only channel
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan int)

	go sendData(ch)
	receiveData(ch)
}
