package learn

import (
	list3 "container/list"
	"fmt"
	"sync"
)

// 将int声明为ChipType类型(将ChipType定义为int类型)
type ChipType int

// 定义类型别名
type intAlias = int

//数组
var list [3]int
var list1 = [...]int{1, 2, 3, 4}
var list2 = [...]int{1, 2, 3, 4}

var array [3][2]int

const (
	None ChipType = iota
	CPU
	GPU
)

//当ChipType类型需要显示为字符串时，Go 语言会自动寻找 String() 方法并进行调用
//不能在一个非本地的类型上定义新方法(使用处和定义处所在包不同)
func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

func Learn03() {
	fmt.Printf("%s %d\n", CPU, CPU)
	var alias intAlias
	alias = 99
	fmt.Println(alias)
	// 打印索引和元素
	for i, v := range list {
		fmt.Printf("%d %d\n", i, v)
	}
	for i, v := range list1 {
		fmt.Printf("%d %d\n", i, v)
	}
	fmt.Println(list1 == list2)
	array = [3][2]int{0: {1, 2}, 2: {6, 7}}
	fmt.Println(list2[2:4])
	fmt.Println(list2[1:])
	// 声明整型切片
	var numList [2]int
	fmt.Println(numList)
	//预分配 2 个元素的切片，只是 b 的内部存储空间已经分配了 5 个，但实际使用了 2 个元素
	//容量不会影响当前的元素个数
	ss := make([]int, 2, 6)
	fmt.Printf("%v - %v \n", ss, len(ss))
	//从开头追加
	newSS := append([]int{11, 12, 13}, ss...)
	fmt.Printf("%v - %v - %v\n", newSS, len(newSS), cap(newSS))
	//从结尾追加(性能好于从开头加)
	newSS1 := append(ss, 99, 98)
	fmt.Printf("%v - %v - %v\n", newSS1, len(newSS1), cap(newSS1))
	copy(newSS, newSS1)
	fmt.Printf("copy %v - %v - %v\n", newSS, len(newSS), cap(newSS))
	//test
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Println(slice2)
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1)
	//并发安全map,有性能损失
	var scene sync.Map
	scene.Store("loren", "cool")
	scene.Store("loren1", "ugly")
	scene.Store("loren2", "very cool")
	scene.Delete("loren1")
	scene.Range(func(key, value interface{}) bool {
		fmt.Printf("%v - %v\n", key, value)
		return true
	})
	//list
	li := list3.New()
	li.PushBack("loren")
	element := li.PushFront(18)
	li.InsertAfter(false, element)
	fmt.Println(li.Len())
	for i := li.Front(); i != nil; i = i.Next() {
		fmt.Printf("%v", i.Value)
	}

}
