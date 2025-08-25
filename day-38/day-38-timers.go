package main

import (
	"fmt"
	"time"
)

func main() {
	// 2 second ka timer banate hain
	timer1 := time.NewTimer(2 * time.Second)

	fmt.Println("Timer set kiya... waiting ⏳")

	// Wait until the timer channel sends a signal
	<-timer1.C
	fmt.Println("Timer 1 fired! 🎉")

	// Timer ko cancel bhi kar sakte ho
	timer2 := time.NewTimer(3 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired! 🚀")
	}()

	stop := timer2.Stop()
	if stop {
		fmt.Println("Timer 2 cancelled ❌")
	}
}
