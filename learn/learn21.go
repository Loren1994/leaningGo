package learn

import (
	"fmt"
	"sync"
)

var (
	count int
	//互斥锁
	countGuard sync.Mutex
	//读写互斥锁
	countGuardRW sync.RWMutex
)

func Learn21() {
	for i := 0; i < 10; i++ {
		go func() {
			setCount(i)
			//fmt.Println(getCount())
			fmt.Println(getRWCount())
		}()
	}
}

//在读多写少的环境中，可以优先使用读写互斥锁（sync.RWMutex），它比互斥锁更加高效
func getRWCount() int {
	//此时另外一个 goroutine 并发访问了 countGuard，
	// 同时也调用了 countGuard.RLock() 时，并不会发生阻塞
	countGuardRW.RLock()
	defer countGuardRW.RUnlock()
	return count
}

func getCount() int {
	//另外一个 goroutine 尝试继续加锁时将会发生阻塞，直到这个 countGuard 被解锁。
	countGuard.Lock()
	defer countGuard.Unlock()
	return count
}

func setCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}
