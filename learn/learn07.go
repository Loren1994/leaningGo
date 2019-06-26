package learn

import "fmt"

// 闭包: http://c.biancheng.net/view/59.html

type Invoker interface {
	//调用时会传入一个 interface{} 类型的变量，这种类型的变量表示任意类型的值
	Call(interface{})
}

type Struct struct {
}

// (s Struct)也可
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

func Learn07() {
	//结构体实现接口
	var invoker1 Invoker = &Struct{} //new(Struct)
	invoker1.Call("hello interface1")
	//函数体实现接口
	var invoker2 Invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})
	invoker2.Call("hello interface2")
	//在闭包内部修改引用的变量
	str := "hello"
	foo := func() {
		str = "world"
	}
	foo()
	fmt.Println(str)
	//累加器
	accumulate1 := accumulate(1)
	fmt.Println(accumulate1())
	fmt.Println(accumulate1())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulate1)
	accumulator2 := accumulate(10)
	fmt.Println(accumulator2())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator2)
	//defer
	fmt.Println("defer begin")
	// 将defer放入延迟调用栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	// 最后一个放入, 位于栈顶, 最先调用
	defer fmt.Println(3)
	fmt.Println("defer end")
}

//函数定义为类型
type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}) {
	f(p)
	//fmt.Println(p)
}

func accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}
