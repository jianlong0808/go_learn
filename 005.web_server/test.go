package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var wgMy sync.WaitGroup

func main() {
	//并发数
	syncNum := 100
	//异步并发的访问
	for i := 0; i < syncNum; i++ {
		wgMy.Add(1)
		//匿名函数扔给协程调度器
		go func(i int) {
			reps, err := http.Get("http://localhost:8080/name")
			if err == nil {
				fmt.Fprint(os.Stdout, strconv.Itoa(i)+"\t"+reps.Status+"\n")
			} else {
				fmt.Println("error")
			}
			//通告协程执行完成
			wgMy.Done()
		}(i)

	}
	//等待所有协程执行完成
	wgMy.Wait()
}
