package learn

import "fmt"

func Learn14() {
	c := make(chan interface{})
	//并发执行
	go printer(c)
	for i := 1; i < 10; i++ {
		c <- i
	}
	//通知结束接收
	c <- 0
	//等待printer结束
	res := <-c
	fmt.Println("main结束", res)
}

func printer(c chan interface{}) {
	//无线循环接收
	for {
		data := <-c
		//将0视为结束标志位
		if data == 0 {
			break
		}
		fmt.Println(data)
	}
	//通知main结束
	c <- 0
}
