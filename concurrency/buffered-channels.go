package concurrency

import "fmt"

func TourBufferedChannels() {
	ch := make(chan int, 2)
	ch <- 1
	fmt.Println(<-ch)
	ch <- 2
	fmt.Println(<-ch)
	ch <- 3
	fmt.Println(<-ch)
	// fmt.Println(<-ch) // empty buffer -> block -> fatal error: all goroutines are asleep - deadlock!

	ch <- 1
	ch <- 2
	// ch <- 3 // full buffer -> block -> fatal error: all goroutines are asleep - deadlock!
}
