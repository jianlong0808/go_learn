package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func fetchAll() {
	start := time.Now()
	//声明一个非阻塞式channel
	ch := make(chan string, 3)
	var wg sync.WaitGroup
	for _, url := range os.Args[1:] {
		wg.Add(1)
		//加入协程, 也就是扔给协程调度器(go 关键字控制)
		go fetch(url, ch)
		wg.Done()
	}
	//等待所有协程结束
	wg.Wait()

	for range os.Args[1:] {
		//从channel中获取值
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("While reding url %s: %v", url, err)
		return
	}

	numBytes, er := io.Copy(io.Discard, resp.Body) //io.Discard类似于linux里的/dev/null
	//防止资源泄露
	resp.Body.Close()
	if er != nil {
		ch <- fmt.Sprintf("While reding url %s: %v", url, er)
		return
	}

	ch <- fmt.Sprintf("%.2fs\t%d\t%s", time.Since(start).Seconds(), numBytes, url)
}

func main() {
	fetchAll()
}
