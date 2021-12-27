package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Increment(key string) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	c.v[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	return c.v[key]
}

func TourMutexCounter() {
	// 1억 tps 테스트. -> 43초 소비. mutex는 상당히 문제있다.
	startTime := time.Now()
	fmt.Println("start time:", startTime)
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 100000000; i++ {
		go c.Increment("key1")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("key1"))
	endTime := time.Now()
	fmt.Println("end time:", time.Now())
	fmt.Println("spend time:", endTime.Sub(startTime))
}
