package concurrency

import "fmt"

func fibonacci2(c, quit chan int) {
	cur, next := 0, 1
	for {
		select { // case를 수행함에 select에서 block해줌. case가 모두 준비됐으면, 무작위로 select됨.
		case c <- cur:
			cur, next = next, cur+next
		case <-quit: // 별도로 quit기능을 위한 channel
			fmt.Println("quit requested.")
			return
		}
	}
}

func TourSelect() {
	c, quit := make(chan int), make(chan int)
	go func() { // 별도 goroutine에서 순차적으로 c에서 10개 가져오고, quit을 발송하므로 순서상 문제 없다. fibonacci2 가 실행되면서 select에서 block상태였을 것.
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) //1 ~ 10 print
		}
		quit <- 0 // end
	}()
	fibonacci2(c, quit)
}
