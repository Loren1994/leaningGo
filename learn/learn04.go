package learn

import "fmt"

func Learn04() {
	var sum = 0
	//1
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("%v\n", sum)
	//2
	sum = 0
	for {
		sum++
		if sum > 100 {
			break
		}
	}
	fmt.Printf("%v", sum)
	//3
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 && i > 1 {
				goto OutLoop
			}
			fmt.Printf("%v - %v\n", i, j)
		}
	}
OutLoop:
	{
		fmt.Println(">>>OutLoop")
		for i := 1; i <= 9; i++ {
			for j := 1; j <= i; j++ {
				fmt.Printf(" %d x %d = %d ", i, j, i*j)
			}
			fmt.Println()
		}
	}
	//4
	for k, v := range []int{1, 2, 3, 4} {
		fmt.Printf("%d - %d\n", k, v)
	}
	//5
	var a = "hello"
	switch a {
	case "hello", "hello1":
		fmt.Println(1)
		fallthrough //紧接着下一个 case 执行,不推荐使用
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}
	//6
	switch {
	case len(a) > 0 && len(a) < 8:
		fmt.Println(3)
	}
	//7
	var t interface{}
	t = &sum //true//"this is string"
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T", t)
	case bool:
		fmt.Printf("type: %t\n", t)
	case string:
		fmt.Printf("type: %T\n", t)
	case *int:
		fmt.Printf("type: %T\n", t) // t 的类型是 int
	case *bool:
		fmt.Printf("type: %t\n", *t) // t 的类型是 *bool
	}

}
