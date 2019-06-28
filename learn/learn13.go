package learn

import (
	"fmt"
	"time"
)

func Learn13() {
	//通道是引用类型，需要使用 make 进行创建
	channel := make(chan interface{})
	go func() {
		time.Sleep(time.Second * 3)
		channel <- "channel message"
		time.Sleep(time.Second * 3)
		channel <- 999
	}()
	//阻塞接收数据
	//data := <-channel
	//fmt.Println(data)
	//非阻塞接收数据,可能造成高的 CPU 占用，因此使用非常少
	//data, ok := <-channel
	//循环接收
	for data := range channel {
		fmt.Println(data)
	}
	//由于接收 goroutine 已经退出，没有 goroutine 发送到通道，因此运行时将会触发宕机报错。
	fmt.Println(">>>>>>end<<<<<<")
}
