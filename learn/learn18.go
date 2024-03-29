package learn

import (
	"fmt"
	"time"
)

func Learn18() {
	// 声明一个退出用的通道
	exit := make(chan int)
	// 打印开始
	fmt.Println("start")
	// 过1秒后, 调用匿名函数
	time.AfterFunc(time.Second, func() {
		// 1秒后, 打印结果
		fmt.Println("one second after")
		// 通知main()的goroutine已经结束
		exit <- 0
	})
	// 等待结束
	<-exit
	//定时
	ticker := time.NewTicker(time.Second)
	timer := time.NewTimer(time.Second * 3)
	var i int
	for {
		select {
		case <-timer.C:
			fmt.Println("timer one")
			goto Exit
		case <-ticker.C:
			i++
			fmt.Println("tick:", i)
		}
	}
Exit:
	fmt.Println("complete!")
}
