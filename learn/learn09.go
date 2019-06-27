package learn

import "fmt"

type Test struct {
	x int
	y int
}

// 小对象由于值复制时的速度较快，所以适合使用非指针接收器。大对象因为复制性能较低，适合使用指针接收器
// 在接收器和参数间传递时不进行复制，只是传递指针。
func Learn09() {
	var t Test
	t.x = 10
	t.y = 9
	ins := &t
	ins1 := &Test{}
	ins1.y = 2
	ins1.x = 4
	fmt.Printf("%+v - %p - %p - %+v\n", t, ins, ins1, *ins1)
	t.setX(999)
	fmt.Printf("%+v - %p\n", t, &t.x)
	t.setValue(888)
	fmt.Printf("%+v - %p\n", t, &t.x)
}

// 指针类型接收器
// 由于指针的特性，调用方法时，修改接收器指针的任意成员变量，在方法结束后，修改都是有效的
func (t *Test) setX(newX int) {
	t.x = newX
	fmt.Println(t.x, &t.x)
}

// 非指针类型接收器
// 当方法作用于非指针接收器时，Go会在代码运行时将接收器的值复制一份
// 在非指针接收器的方法中可以获取接收器的成员值，但修改后无效
func (t Test) setValue(newX int) {
	t.x = newX
	fmt.Println(t.x, &t.x)
}
