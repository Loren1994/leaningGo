package learn

import (
	"fmt"
	"net/http"
	"sync"
)

func Learn22() {
	var urls = []string{
		"http://www.github.com/",
		"https://www.qiniu.com/",
		"https://www.golangtc.com/",
	}
	//等待组
	var wg sync.WaitGroup
	for _, url := range urls {
		//开启一个任务,等待组就增加1
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			_, error := http.Get(u)
			fmt.Println(u, error)
		}(url)
	}
	wg.Wait()
	fmt.Println("over")
}
