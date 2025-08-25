package main

import "fmt"

func main() {
	queue := make(chan string, 3)
	queue <- "task1"
	queue <- "task2"
	queue <- "task3"
	close(queue)

	// range loop over channel
	for elem := range queue {
		fmt.Println(elem)
	}
}
