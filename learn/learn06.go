package learn

import (
	"fmt"
	"strings"
)

func Learn06() {
	//1
	var ff func()
	ff = printTest
	ff()
	//2
	list := []string{" aaa ", "bbb  ", "   ccc"}
	chain := []func(string) string{
		strings.TrimSpace,
		strings.ToUpper,
	}
	stringProcess(list, chain)
	//3
	visit([]string{"aaa", "bbb", "ccc"}, func(s string) {
		fmt.Printf("%v ", s)
	})
}

func visit(list []string, callback func(string)) {
	for _, v := range list {
		callback(v)
	}
}

func stringProcess(list []string, chain []func(string) string) {
	for index, str := range list {
		result := str
		for _, ch := range chain {
			result = ch(result)
		}
		fmt.Printf("%v - %v\n", index, result)
	}
}

func printTest() {
	fmt.Println("this is printTest()")
}
