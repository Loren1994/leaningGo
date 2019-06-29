package learn

import "fmt"

func Learn16() {
	//缓冲大小为5个元素的channel
	ch := make(chan interface{}, 3)
	//查看当前通道的大小
	fmt.Println(len(ch))
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(len(ch))
	//ch <- 4
}
