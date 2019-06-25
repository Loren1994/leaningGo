package main

import (
	"flag"
	"fmt"
)

// 定义一个 mode 变量，这个变量的类型是 *string
var mode = flag.String("mode", "default value", "mode usage")

func main() {
	//指针取值
	field1 := "test"
	ptr := &field1
	fmt.Printf("%v %T\n", ptr, ptr)
	value := *ptr
	fmt.Printf("%v %T\n\n", value, value)
	//指针交换值
	x, y := 1, 2
	//swap1
	fmt.Printf("交换前:%v - %v - %v - %v\n", x, y, &x, &y)
	swap1(&x, &y)
	fmt.Printf("交换后:%v - %v - %v - %v\n\n", x, y, &x, &y)
	//swap2
	fmt.Printf("交换前:%v - %v - %v - %v\n", x, y, &x, &y)
	swap2(&x, &y)
	fmt.Printf("交换后:%v - %v - %v - %v\n\n", x, y, &x, &y)
	//获取命令行参数
	//解析参数
	flag.Parse()
	fmt.Println(*mode)
	fmt.Printf("mode type : %T \n\n", mode)
	//创建指针
	str := new(string)
	*str = "pointer value"
	fmt.Printf(*str)
}

// b, a = a, b 指针的交换操作不成功
// 原因为:
// 指针a和b的变量值确实被交换,但和a、b关联的两个变量(x,y)并没有实际关联
// 即a,b交换成功, 但x,y交换失败
// 比如写有两座房子的卡片放在桌上一字摊开，交换两座房子的卡片后并不会对两座房子有任何影响
func swap1(a, b *int) {
	fmt.Printf("swap1前:%v - %v - %v - %v\n", a, b, *a, *b)
	a, b = b, a
	fmt.Printf("swap1后:%v - %v - %v - %v\n", a, b, *a, *b)
}

//使用中间变量交换值
func swap2(a, b *int) {
	fmt.Printf("swap2前:%v - %v - %v - %v\n", a, b, *a, *b)
	t := *a
	*a = *b
	*b = t
	fmt.Printf("swap2后:%v - %v - %v - %v\n", a, b, *a, *b)
}
