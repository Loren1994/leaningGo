package learn

import (
	"fmt"
	"time"
)

//select 多路复用
func Learn17() {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	go chReceiver(ch2)

	//ch1 <- "message" //向ch1发送
	//res1 := executor(ch1, ch2)
	//fmt.Println(res1)
	ch2 <- "message" //向ch2发送
	//close(ch2)
	res2 := executor(ch1, ch2)
	fmt.Println(res2)

	ch2 <- "message" //向ch2发送
	fmt.Println(executor(ch1, ch2))
}

func chReceiver(ch chan interface{}) {
	for {
		data := <-ch
		fmt.Println("接收了", data)
		//模拟超时
		//time.Sleep(time.Second * 4)
		ch <- "copy that"
	}
}

func executor(ch1 chan interface{}, ch2 chan interface{}) interface{} {
	//多个操作在每次 select 中挑选一个进行响应
	var res interface{}
	select {
	case p := <-ch1: //ch1接收
		res = p
	case ch2 <- "case message": //向ch2发送
		res = "ch2 send success"
	case p := <-ch2: // ch2返回
		res = p
	case <-time.After(time.Second): // 超时
		res = "Time out"
	}
	return res
}
