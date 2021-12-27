package concurrency

// 주의: 절대로 수신자가 아닌 전송자만이 channel을 닫아야합니다. 닫힌 channel에 전송하는 것은 panic을 야기할 것입니다.
// 하나 더 주의: Channel은 파일과 다릅니다. 당신은 file과 달리 보통 channel을 닫을 필요는 없습니다. channel을 닫는 것은 range 반복문을 종료시키는 것과 같이 수신자가 더 이상 들어오는 값이 없다는 것을 알아야하는 경우에만 필요합니다.

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	cur, next := 0, 1
	for i := 0; i < n; i++ {
		c <- cur
		cur, next = next, cur+next
	}
	close(c)
}

func TourRangeAndClose() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for v := range c {
		fmt.Println(v)
	}
}
