package concurrency

import (
	"fmt"
	"time"
)

func say(s *string) {
	fmt.Println(s, *s)
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(*s)
	}
}

func TourGoroutines() {
	s := "world"
	go say(&s)
	go say(&s)
	s2 := "hello"
	say(&s2)
}
