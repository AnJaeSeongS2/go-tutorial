package concurrency

// select 절에서 default case는 다른 case가 준비되지 않았을 때 실행됨.

import (
	"fmt"
	"time"
)

func TourDefaultSelection() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case t := <-tick:
			fmt.Println("tick:", t)
		case t := <-boom:
			fmt.Println("boom:", t)
			return
		default:
			fmt.Println(".....")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
