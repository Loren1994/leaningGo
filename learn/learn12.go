package learn

import (
	"fmt"
	"runtime"
	"time"
)

func Learn12() {
	fmt.Println("sleep start...")
	// 并发
	// go创建goroutine,会忽略返回值
	go timeFun()
	fmt.Println("sleep end...")
	// 接受命令行输入, 不做任何事情
	//var input string
	//i, err := fmt.Scanln(&input)
	//fmt.Println(i, err)
	go func() {
		time.Sleep(time.Second * 5)
	}()
	fmt.Println("func sleep end...")
	//设置逻辑CPU数量
	//在 GOMAXPROCS 数量与任务数量相等时，可以做到并行执行，但一般情况下都是并发执行
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("逻辑CPU数量:", runtime.NumCPU())
}

func timeFun() {
	time.Sleep(time.Second * 5)
}
