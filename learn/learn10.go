package learn

import "fmt"

var eventByName = make(map[string][]func(interface{}))

//注册
func registerEvent(name string, callback func(interface{})) {
	list := eventByName[name]
	list = append(list, callback)
	eventByName[name] = list
}

//调用
func callEvent(name string, param interface{}) {
	list := eventByName[name]
	for _, callback := range list {
		callback(param)
	}
}

func Learn10() {
	registerEvent("loren", func(i interface{}) {
		fmt.Println("收到loren消息:", i)
	})
	registerEvent("test", func(i interface{}) {
		fmt.Println("收到test消息:", i)
	})
	callEvent("loren", "send message")
	callEvent("test", "test message")
}
