package learn

import "fmt"

type Data struct {
	complax  []int
	instance InnerData
	ptr      *InnerData
}

type InnerData struct {
	a int
}

func passByValue(inFunc Data) Data {
	fmt.Printf("%+v - %p\n", inFunc, &inFunc)
	return inFunc
}

func Learn05() {
	//1
	i, s := test(6, 6)
	fmt.Printf("%v %v\n", i, s)
	err, suc := test1(1, 1)
	fmt.Printf("%v %v\n", err, suc)
	err2, suc2 := test2(1, 1)
	fmt.Printf("%v %v\n", err2, suc2)
	//2
	d, h, m := convertTime(18000)
	fmt.Printf("%v - %v - %v\n", d, h, m)
	//3
	in := Data{
		complax:  []int{1, 2, 3},
		instance: InnerData{5},
		ptr:      &InnerData{1},
	}
	// 输入结构的成员情况
	fmt.Printf("in value: %+v\n", in)
	// 输入结构的指针地址
	fmt.Printf("in ptr: %p\n", &in)
	// 传入结构体，返回同类型的结构体
	out := passByValue(in)
	// 输出结构的成员情况
	fmt.Printf("out value: %+v\n", out)
	// 输出结构的指针地址
	fmt.Printf("out ptr: %p\n", &out)
	// 结论:
	// 所有的 Data 结构的指针地址发生了变化，意味着所有的结构都是一块新的内存，
	// 无论是将 Data 结构传入函数内部，还是通过函数返回值传回 Data 都会发生复制行为。
	// 所有的 Data 结构中的成员值都没有发生变化，原样传递，意味着所有参数都是值传递。
	// Data 结构的 ptr 成员在传递过程中保持一致，
	// 表示指针在函数参数值传递中传递的只是指针值，不会复制指针指向的部分。
}

func convertTime(seconds int) (day int, hour int, minute int) {
	day = seconds / (60 * 60 * 24)
	hour = seconds / (60 * 60)
	minute = seconds / 60
	return
}

func test(a, b int) (string, int) {
	return "print:", a + b
}

func test1(a, b int) (err string, success int) {
	return "sss", 1
}

func test2(a, b int) (err, success int) {
	return 1, 1
}
