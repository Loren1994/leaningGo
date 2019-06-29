package learn

import (
	"fmt"
	"sync/atomic"
)

var seq int64

func Learn20() {
	for i := 0; i < 10; i++ {
		//生成10个并发序列号
		go GenerateID()
	}
	fmt.Println(GenerateID())
}

func GenerateID() int64 {
	//存在竞态问题
	//atomic.AddInt64(&seq, 1)
	//return seq
	//原子的增加序列号,不存在竞态
	return atomic.AddInt64(&seq, 1)
}
