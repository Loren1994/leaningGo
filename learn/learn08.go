package learn

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	valueByKey = make(map[string]int)
	// 保证使用映射时的并发安全的互斥锁
	valueByKeyGuard sync.Mutex
)

func Learn08() {
	start := time.Now()
	getValue()
	getValueDefer()
	end := time.Now()
	fmt.Println(end.Sub(start))
	//down
	//defer fmt.Println("宕机后要做的事情1")
	//panic("make crash")
	// 故意造成空指针访问错误
	ProtectRun(func() {
		fmt.Println("赋值宕机前")
		var a *int
		*a = 1
		fmt.Println("赋值宕机后")
	})
}

// 保护方式允许一个函数
func ProtectRun(entry func()) {
	// 延迟处理的函数
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error:", err)
		default: // 非运行时错误
			fmt.Println("error:", err)
		}
	}()
	entry()
}

func getValue() int {
	valueByKeyGuard.Lock()
	v := valueByKey[""]
	valueByKeyGuard.Unlock()
	return v
}

//defer简化
func getValueDefer() int {
	valueByKeyGuard.Lock()
	//延迟到函数结束时调用
	defer valueByKeyGuard.Unlock()
	return valueByKey[""]
}
